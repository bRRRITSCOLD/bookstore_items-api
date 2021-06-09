package ping_controllers

import (
	"net/http"
)

var (
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type pingController struct{}

func (pc *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
