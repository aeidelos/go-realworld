package user

import (
    "github.com/neotroops/go-realworld/components"
    "github.com/neotroops/go-realworld/configs"
    "github.com/neotroops/go-realworld/models"
)

type ServiceContract interface {
    Register(form models.UserRegisterForm) models.UserResponse
    Login(form models.UserLogin) models.UserResponse
    SaveToken(form models.UserLogin)
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

func (s Service) Login(form models.UserLogin) models.UserResponse {
    return models.UserResponse{
        Username: form.Username,
        Email: form.Password,
    }
}

func (s Service) SaveToken(form models.UserLogin) {
    db := components.GetConnection(configs.AllConfig().Db)
    db.Query("UPDATE user SET token=? where username=?",form.Token, form.Username)

}
