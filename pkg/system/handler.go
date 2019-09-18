package system

import (
    "github.com/sirupsen/logrus"
    "net/http"
)

type HandlerContract interface {
    DefaultHandler(w http.ResponseWriter, r *http.Request)
    PingHandler(w http.ResponseWriter, r *http.Request)
}

type Handler struct {}

func NewHandler() *Handler {
    return &Handler{}
}

func (h *Handler) DefaultHandler(w http.ResponseWriter, r *http.Request) {
    defaultResponseWriter(w, "Welcome to Neo Application")
}

func (h *Handler) PingHandler(w http.ResponseWriter, r *http.Request) {
    defaultResponseWriter(w, "pong")
}

func defaultResponseWriter(w http.ResponseWriter, value string) {
    if _, err := w.Write([]byte(value)); err != nil {
        logrus.Panic("can't write response to ")
    }
}
