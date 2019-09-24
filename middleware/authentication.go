package middleware

import (
    "github.com/neotroops/go-realworld/components"
    "github.com/neotroops/go-realworld/configs"
    "github.com/neotroops/go-realworld/models"
    "github.com/sirupsen/logrus"
    "log"
    "net/http"
)

type AuthenticationMiddleware struct {}

func (amw *AuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        token := r.Header.Get("Authorization")
        logrus.Info(token)
        db := components.GetConnection(configs.AllConfig().Db)
        var user = models.UserResponse{}
        err := db.Get(&user,"SELECT * FROM user where token=?",token)
        if err == nil {
            log.Printf("Authenticated user %s\n", user)
            next.ServeHTTP(w, r)
        } else {
            panic(err)
            http.Error(w, "Forbidden", http.StatusForbidden)
        }
    })
}
