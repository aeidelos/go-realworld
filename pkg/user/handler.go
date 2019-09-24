package user

import (
    "encoding/json"
    "github.com/neotroops/go-realworld/models"
    "github.com/sirupsen/logrus"
    "golang.org/x/crypto/bcrypt"
    "net/http"
)

type Handler struct {
    Service ServiceContract
}

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
    var user models.UserRegisterForm
    err := json.NewDecoder(r.Body).Decode(&user)
    logrus.Info(h.Service.Register(user), err)
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
    var user models.UserLogin
    err := json.NewDecoder(r.Body).Decode(&user)
    logrus.Info(h.Service.Login(user), err)

    hash, err := bcrypt.GenerateFromPassword([]byte(user.Username), bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
    user.Token = string(hash)
    userJson, err := json.Marshal(user)
    if err != nil {
        panic(nil)
    }
    h.Service.SaveToken(user)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(userJson)
}
