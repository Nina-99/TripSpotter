package service

import (
	"github.com/Nina-99/TripSpotter/backend/data/request"
	"github.com/Nina-99/TripSpotter/backend/models"
)

type UserService interface {
	Login(req request.LoginUserRequest) (string, error)
	Register(req request.RegisterUserRequest) (string, error)
	Update(userId uint, req request.RegisterUserRequest) error
	Delete(userId uint) error
	FindAll() ([]models.User, error)
}
