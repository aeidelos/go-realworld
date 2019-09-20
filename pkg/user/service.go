package user

import "github.com/neotroops/go-realworld/models"

type ServiceContract interface {
    Register(form models.UserRegisterForm) models.UserResponse
}

type Service struct {}

func NewService() Service {
    return Service{}
}

func (s Service) Register(form models.UserRegisterForm) models.UserResponse {
    return models.UserResponse{
        Username: form.Username,
        Email: form.Email,
    }
}

