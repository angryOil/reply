package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"reply/cmd/app/handler"
	"reply/internal/controller"
	"reply/internal/repository"
	"reply/internal/repository/infla"
	"reply/internal/service"
)

func main() {
	r := mux.NewRouter()
	h := getHandler()
	r.PathPrefix("/replies").Handler(h)

	err := http.ListenAndServe(":8090", r)
	if err != nil {
		panic(err)
	}
}

func getHandler() http.Handler {
	return handler.NewHandler(controller.NewController(service.NewService(repository.NewRepository(infla.NewDB()))))
}
