package service

import (
	"errors"

	"github.com/Nina-99/TripSpotter/backend/config"
	"github.com/Nina-99/TripSpotter/backend/data/request"
	"github.com/Nina-99/TripSpotter/backend/helper"
	"github.com/Nina-99/TripSpotter/backend/models"
	"github.com/Nina-99/TripSpotter/backend/repository"
	"github.com/Nina-99/TripSpotter/backend/utils"
	"github.com/go-playground/validator/v10"
)

type AuthenticationServiceImpl struct {
	UsersRepository repository.UsersRepository
	Validate        *validator.Validate
}

func NewAuthenticationServiceImpl(usersRepository repository.UsersRepository, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		UsersRepository: usersRepository,
		Validate:        validate,
	}
}

// Login implements AuthenticationService
func (a *AuthenticationServiceImpl) Login(users request.LoginUserRequest) (string, error) {
	// Find username in database
	new_users, users_err := a.UsersRepository.FindByUsername(users.Username)
	if users_err != nil {
		return "", errors.New("invalid username or Password")
	}

	config, _ := config.LoadConfig(".")

	verify_error := utils.VerifyPassword(new_users.Password, users.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or Password")
	}

	// Generate Token
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, new_users.Id, config.TokenSecret)
	helper.ErrorPanic(err_token)
	return token, nil

}

// Register implements AuthenticationService
func (a *AuthenticationServiceImpl) Register(users request.CreateUserRequest) {

	hashedPassword, err := utils.HashPassword(users.Password)
	helper.ErrorPanic(err)

	newUser := models.User{
		Username: users.Username,
		Email:    users.Email,
		Password: hashedPassword,
		Role:     users.Role,
	}
	a.UsersRepository.Save(newUser)
}
