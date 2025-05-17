package service

import "github.com/Nina-99/TripSpotter/backend/data/request"

type AuthenticationService interface {
	Login(user request.LoginUserRequest) (string, error)
	Register(user request.CreateUserRequest)
}
