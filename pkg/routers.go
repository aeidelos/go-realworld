package pkg

import (
    "context"
    "github.com/gorilla/mux"
    negronilogrus "github.com/meatballhat/negroni-logrus"
    "github.com/neotroops/go-realworld/configs"
    "github.com/neotroops/go-realworld/middleware"
    "github.com/neotroops/go-realworld/pkg/system"
    "github.com/sirupsen/logrus"
    "github.com/urfave/negroni"
    "net/http"
    "os"
    "time"
)

func StartAPIServer(config configs.App, s <-chan os.Signal) {
    var wait time.Duration
    wait = time.Second * 1
    r := mux.NewRouter()
    n := negroni.New(negroni.NewRecovery(), negronilogrus.NewMiddleware())
    services := NewServices()
    defaultRouter(r)
    userRouter(r, services)
    n.UseHandler(r)
    srv := &http.Server{
        Handler: n,
        Addr: ":" + config.AppPort,
        WriteTimeout: 15 * time.Second,
        ReadTimeout: 15 * time.Second,
    }
    go func() {
        <-s
        ctx, cancel := context.WithTimeout(context.Background(), wait)
        defer cancel()
        if err := srv.Shutdown(ctx); err != nil {
            logrus.Panic(err)
            os.Exit(1)
        }
        logrus.Info("shutdown server")
        os.Exit(0)
    }()
    logrus.Fatal(srv.ListenAndServe())
}

func defaultRouter (r *mux.Router) {
    amw := middleware.AuthenticationMiddleware{}
    defaultHandler := system.NewDefaultHandler()
    s := r.PathPrefix("/").Subrouter()
    s.HandleFunc("/", defaultHandler.RootHandler)
    s.HandleFunc("/ping", defaultHandler.PingHandler)
    s.Use(amw.Middleware)
}

func userRouter (r *mux.Router, services ServiceInitializerContract) {
    handler := NewUserHandler(services)
    r.HandleFunc("/api/users", handler.Register).Methods(http.MethodPost)
    r.HandleFunc("/login", handler.Login).Methods(http.MethodPost)
}
