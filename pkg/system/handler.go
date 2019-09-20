 package system

import (
    "github.com/sirupsen/logrus"
    "net/http"
)

type DefaultHandlerContract interface {
    RootHandler(w http.ResponseWriter, r *http.Request)
    PingHandler(w http.ResponseWriter, r *http.Request)
}

type DefaultHandler struct {}

func NewDefaultHandler() *DefaultHandler {
    return &DefaultHandler{}
}

func (h *DefaultHandler) RootHandler(w http.ResponseWriter, r *http.Request) {
    defaultResponseWriter(w, "Welcome to Neo Application")
}

func (h *DefaultHandler) PingHandler(w http.ResponseWriter, r *http.Request) {
    defaultResponseWriter(w, "pong")
}

func defaultResponseWriter(w http.ResponseWriter, value string) {
    if _, err := w.Write([]byte(value)); err != nil {
        logrus.Panic("can't write response to ")
    }
}
