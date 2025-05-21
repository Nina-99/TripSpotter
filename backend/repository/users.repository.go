package repository

import "github.com/Nina-99/TripSpotter/backend/models"

type UsersRepository interface {
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(user *models.User) error
	FindById(userId uint) (*models.User, error)
	FindAll() ([]models.User, error)
	FindByEmail(email string) (*models.User, error)
}
