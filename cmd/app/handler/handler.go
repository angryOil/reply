package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"reply/internal/controller"
	"reply/internal/controller/req"
	"strconv"
)

type Handler struct {
	c controller.Controller
}

func NewHandler(c controller.Controller) http.Handler {
	r := mux.NewRouter()
	h := Handler{c: c}
	r.HandleFunc("/replies/{cafeId:[0-9]+}/{boardId:[0-9]+}", h.create).Methods(http.MethodPost)
	return r
}

const (
	InvalidCafeId  = "invalid cafe id"
	InvalidBoardId = "invalid board id"
)

func (h Handler) create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	boardId, err := strconv.Atoi(vars["boardId"])
	if err != nil {
		http.Error(w, InvalidBoardId, http.StatusBadRequest)
		return
	}
	var c req.Create
	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.c.Create(r.Context(), cafeId, boardId, c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
