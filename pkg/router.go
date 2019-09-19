package pkg

import (
    "github.com/gorilla/mux"
    negronilogrus "github.com/meatballhat/negroni-logrus"
    "github.com/neotroops/go-realworld/configs"
    "github.com/neotroops/go-realworld/pkg/system"
    "github.com/sirupsen/logrus"
    "github.com/urfave/negroni"
    "net/http"
    "time"
)

func StartAPIServer(config configs.App) {
    r := mux.NewRouter()
    n := negroni.New(negroni.NewRecovery(), negronilogrus.NewMiddleware())
    defaultRouter(r)
    n.UseHandler(r)
    srv := &http.Server{
        Handler: n,
        Addr: ":" + config.AppPort,
        WriteTimeout: 15 * time.Second,
        ReadTimeout: 15 * time.Second,
    }
    logrus.Fatal(srv.ListenAndServe())
}

func defaultRouter (r *mux.Router) {
    defaultHandler := system.NewHandler()
    r.HandleFunc("/", defaultHandler.DefaultHandler)
    r.HandleFunc("/ping", defaultHandler.PingHandler)
}
