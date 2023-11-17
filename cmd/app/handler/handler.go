package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"reply/internal/controller"
)

type Handler struct {
	c controller.Controller
}

func NewHandler(c controller.Controller) http.Handler {
	r := mux.NewRouter()
	//h := Handler{c: c}

	return r
}
