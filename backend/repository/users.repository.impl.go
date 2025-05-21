package repository

import (
	"github.com/Nina-99/TripSpotter/backend/models"
	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	db *gorm.DB
}

func NewUsersRepositoryImpl(db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{db: db}
}

// FindAll implements UsersRepository
func (u *UsersRepositoryImpl) FindAll() ([]models.User, error) {
	var users []models.User
	err := u.db.Find(&users).Error
	return users, err
}

// FindById implements UsersRepository
func (u *UsersRepositoryImpl) FindById(userId uint) (*models.User, error) {
	var user models.User
	err := u.db.First(&user, userId).Error
	return &user, err
}

// FindByEmail implements UsersRepository
func (u *UsersRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

// Save implements UsersRepository
func (u *UsersRepositoryImpl) Create(user *models.User) error {
	return u.db.Create(&user).Error
}

// Update implements UsersRepository
func (u *UsersRepositoryImpl) Update(user *models.User) error {
	return u.db.Save(user).Error
}

// Delete implements UsersRepository
func (u *UsersRepositoryImpl) Delete(user *models.User) error {
	return u.db.Delete(user).Error
}
