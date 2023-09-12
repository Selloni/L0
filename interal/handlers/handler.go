package handlers

import (
	"L0/pkg/inmemory"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Handler interface {
	Register(r *httprouter.Router)
}

type handler struct {
	cash *inmemory.Cash
	//order db.Order
}

func NewHandler(cash *inmemory.Cash) Handler {
	return &handler{
		cash: cash,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/", h.GetAllOrders)
	router.GET("/user/:id", h.GetOrder)
}

func (h *handler) GetAllOrders(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	for uid := range h.cash
	if _, err := w.Write([]byte("Get all list")); err != nil {
		log.Fatal(err)
	}
}

func (h *handler) GetOrder(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	w.Write([]byte("Get User\n"))
	w.Write([]byte("Get kk\n"))
	w.Write([]byte("Get oo\n"))
	w.Write([]byte(r.RequestURI))

}

//// методы лежат в модели
//func (h *handler) Register(router *httprouter.Router, o db.Order) {
//	router.GET("/", o.GetAllOrders)
//	router.GET("/user/:id", o.GetOrder)
//}
