package repository

import (
	"errors"

	"github.com/Nina-99/TripSpotter/backend/data/request"
	"github.com/Nina-99/TripSpotter/backend/helper"
	"github.com/Nina-99/TripSpotter/backend/models"
	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsersRepositoryImpl(Db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{Db: Db}
}

// Delete implements UsersRepository
func (u *UsersRepositoryImpl) Delete(usersId int) {
	var users models.User
	result := u.Db.Where("id = ?", usersId).Delete(&users)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UsersRepository
func (u *UsersRepositoryImpl) FindAll() []models.User {
	var users []models.User
	results := u.Db.Find(&users)
	helper.ErrorPanic(results.Error)
	return users
}

// FindById implements UsersRepository
func (u *UsersRepositoryImpl) FindById(usersId int) (models.User, error) {
	var users models.User
	result := u.Db.Find(&users, usersId)
	if result != nil {
		return users, nil
	} else {
		return users, errors.New("users is not found")
	}
}

// Save implements UsersRepository
func (u *UsersRepositoryImpl) Save(users models.User) {
	result := u.Db.Create(&users)
	helper.ErrorPanic(result.Error)
}

// Update implements UsersRepository
func (u *UsersRepositoryImpl) Update(users models.User) {
	var updateUsers = request.UpdateUserRequest{
		Id:       users.Id,
		Username: users.Username,
		Email:    users.Email,
		Password: users.Password,
	}
	result := u.Db.Model(&users).Updates(updateUsers)
	helper.ErrorPanic(result.Error)
}

// FindByUsername implements UsersRepository
func (u *UsersRepositoryImpl) FindByUsername(username string) (models.User, error) {
	var users models.User
	result := u.Db.First(&users, "username = ?", username)

	if result.Error != nil {
		return users, errors.New("invalid username or Password")
	}
	return users, nil
}
