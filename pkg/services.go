package pkg

import "github.com/neotroops/go-realworld/pkg/user"

type ServiceInitializerContract interface {
    UserService() user.ServiceContract
}

type Services struct {
    User user.ServiceContract
}

func NewServices() ServiceInitializerContract{
    var userService user.ServiceContract
    userService = user.NewService()
    return Services{
        User: userService,
    }
}

func (s Services) UserService() user.ServiceContract {
    return s.User
}
