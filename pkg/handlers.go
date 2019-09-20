package pkg

import "github.com/neotroops/go-realworld/pkg/user"

func NewUserHandler(contract ServiceInitializerContract) user.Handler {
    return user.Handler{
        Service: contract.UserService(),
    }
}
