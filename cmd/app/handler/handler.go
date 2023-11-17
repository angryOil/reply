package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"reply/internal/controller"
	"reply/internal/controller/req"
	"reply/internal/controller/res"
	"reply/internal/page"
	"strconv"
	"strings"
)

type Handler struct {
	c controller.Controller
}

func NewHandler(c controller.Controller) http.Handler {
	r := mux.NewRouter()
	h := Handler{c: c}
	r.HandleFunc("/replies/cnt", h.getCountList).Methods(http.MethodGet)
	r.HandleFunc("/replies/{boardId:[0-9]+}", h.getList).Methods(http.MethodGet)
	r.HandleFunc("/replies/{cafeId:[0-9]+}/{boardId:[0-9]+}", h.create).Methods(http.MethodPost)
	r.HandleFunc("/replies/{id}", h.patch).Methods(http.MethodPatch)
	r.HandleFunc("/replies/{id}", h.delete).Methods(http.MethodDelete)
	return r
}

const (
	InternalServerError = "internal server error"
	InvalidId           = "invalid reply id"
	InvalidCafeId       = "invalid cafe id"
	InvalidBoardId      = "invalid board id"
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

func (h Handler) patch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, InvalidId, http.StatusBadRequest)
		return
	}
	var p req.Patch
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.c.Patch(r.Context(), id, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h Handler) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, InvalidId, http.StatusBadRequest)
		return
	}
	err = h.c.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h Handler) getList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	boardId, err := strconv.Atoi(vars["boardId"])
	if err != nil {
		http.Error(w, InvalidBoardId, http.StatusBadRequest)
		return
	}
	reqPage := page.GetPageReqByRequest(r)
	list, total, err := h.c.GetList(r.Context(), boardId, reqPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	listTotalDto := res.ListTotalDto{
		Content: list,
		Total:   total,
	}
	data, err := json.Marshal(listTotalDto)
	if err != nil {
		log.Println("getList json.Marshal err: ", err)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (h Handler) getCountList(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("board-ids")
	if q == "" {
		http.Error(w, InvalidBoardId, http.StatusBadRequest)
		return
	}
	boardIdsArr := stringToIntArr(q)
	cntArr, err := h.c.GetCountList(r.Context(), boardIdsArr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	listDto := res.CountListDto{Content: cntArr}
	data, err := json.Marshal(listDto)
	if err != nil {
		log.Println("getCountList json.Marshal err: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func stringToIntArr(s string) []int {
	s = strings.ReplaceAll(s, " ", "")
	sArr := strings.Split(s, ",")
	intArr := make([]int, 0)
	for _, s := range sArr {
		i, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		intArr = append(intArr, i)
	}
	return intArr
}
