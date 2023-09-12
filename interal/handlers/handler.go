package handlers

import (
	"L0/pkg/inmemory"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type Handler interface {
	Register(r *httprouter.Router)
}

type handler struct {
	cash *inmemory.InMemory
	//order db.Order
}

func NewHandler(cash *inmemory.InMemory) Handler {
	return &handler{
		cash: cash,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/", h.ShowPost)
	router.GET("/uid/", h.ShowPost)
}

//func (h *handler) GetAllOrders(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
//	for key, _ := range *h.cash.GetStore() {
//		if _, err := w.Write([]byte(key)); err != nil {
//			log.Fatal(err)
//		}
//		w.Write([]byte("\n"))
//	}
//}
//
//func (h *handler) GetOrder(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
//	for key, _ := range *h.cash.GetStore() {
//		if _, err := w.Write([]byte(key)); err != nil {
//			log.Fatal(err)
//		}
//		w.Write([]byte("\n"))
//	}
//}

func (h *handler) ShowPost(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	tmpl, err := template.ParseFiles("ui/homePage.html")
	if err != nil {
		return
	}
	tmpl.Execute(w, nil)

}

//// методы лежат в модели
//func (h *handler) Register(router *httprouter.Router, o db.Order) {
//	router.GET("/", o.GetAllOrders)
//	router.GET("/user/:id", o.GetOrder)
//}
