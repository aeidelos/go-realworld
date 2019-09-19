package pkg

import (
    "github.com/gorilla/mux"
    "github.com/neotroops/go-realworld/configs"
    "github.com/neotroops/go-realworld/pkg/system"
    "github.com/sirupsen/logrus"
    "net/http"
    "time"
)

func StartAPIServer(config configs.App) {
    r := mux.NewRouter()
    defaultHandler := system.NewHandler()
    r.HandleFunc("/", defaultHandler.DefaultHandler)
    r.HandleFunc("/ping", defaultHandler.PingHandler)
    srv := &http.Server{
        Handler: r,
        Addr: ":" + config.AppPort,
        WriteTimeout: 15 * time.Second,
        ReadTimeout: 15 * time.Second,
    }
    logrus.Fatal(srv.ListenAndServe())
}
