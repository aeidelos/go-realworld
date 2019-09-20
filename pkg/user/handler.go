package user

import (
    "encoding/json"
    "github.com/neotroops/go-realworld/models"
    "github.com/sirupsen/logrus"
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
