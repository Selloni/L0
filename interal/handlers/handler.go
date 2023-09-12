package handlers

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Handler interface {
	Register(r *httprouter.Router)
}

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/", h.GetAllOrders)
	router.GET("/user/:id", h.GetOrder)
}

func (h *handler) GetAllOrders(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	if _, err := w.Write([]byte("Get all list")); err != nil {
		log.Fatal(err)
	}
}

func (h *handler) GetOrder(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	w.Write([]byte("Get User"))
}
