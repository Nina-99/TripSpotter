package service

import (
	"github.com/Nina-99/TripSpotter/backend/data/request"
	"github.com/Nina-99/TripSpotter/backend/models"
	"github.com/Nina-99/TripSpotter/backend/repository"
	"github.com/Nina-99/TripSpotter/backend/utils"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UsersRepository repository.UsersRepository
	Validate        *validator.Validate
}

func NewUserServiceImpl(usersRepository repository.UsersRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UsersRepository: usersRepository,
		Validate:        validate,
	}
}

// Login implements UserService
func (a *UserServiceImpl) Login(req request.LoginUserRequest) (string, error) {
	// Find username in database
	user, err := a.UsersRepository.FindByEmail(req.Email)
	match := utils.VerifyPassword(req.Password, user.Password)
	if err != nil || match == false {
		return "", err
	}
	return utils.GenerateJWT(user.Id, user.Email)
}

// Register implements UserService
func (a *UserServiceImpl) Register(req request.RegisterUserRequest) (string, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := models.User{Username: req.Username, Email: req.Email, Password: string(hash)}
	if err := a.UsersRepository.Create(&user); err != nil {
		return "", err
	}
	return utils.GenerateJWT(user.Id, user.Email)
}

// Update implements UserServiceImpl
func (a *UserServiceImpl) Update(userId uint, req request.RegisterUserRequest) error {
	user, err := a.UsersRepository.FindById(userId)
	if err != nil {
		return err
	}
	user.Username = req.Username
	user.Email = req.Email
	user.Password = req.Password
	return a.UsersRepository.Update(user)
}

// Delete user implements UserService
func (a *UserServiceImpl) Delete(userId uint) error {
	user, err := a.UsersRepository.FindById(userId)
	if err != nil {
		return err
	}
	return a.UsersRepository.Delete(user)
}

func (a *UserServiceImpl) FindAll() ([]models.User, error) {
	return a.UsersRepository.FindAll()
}
