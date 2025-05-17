package repository

import "github.com/Nina-99/TripSpotter/backend/models"

type UsersRepository interface {
	Save(user models.User)
	Update(user models.User)
	Delete(userId int)
	FindById(userId int) (models.User, error)
	FindAll() []models.User
	FindByUsername(username string) (models.User, error)
}
