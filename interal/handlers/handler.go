package handlers

import (
	"L0/pkg/inmemory"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
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
	router.GET("/uid/", h.ShowOrder)
}

func (h *handler) ShowPost(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	tmpl, err := template.ParseFiles("ui/homePage.html")
	if err != nil {
		return
	}
	tmpl.Execute(w, nil)
}

func (h *handler) ShowOrder(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	find := r.URL.Query().Get("UID")
	cashMap := *h.cash.GetStore()
	if value, ok := cashMap[find]; ok {
		all, err := json.MarshalIndent(value, "", "\t")
		if err != nil {
			log.Printf("failed to wrap in json")
		}
		w.Write(all)
	} else {
		log.Printf("don't expect such a uid =  %s", find)
		w.Write([]byte("don't expect such a uid =" + find))
	}
}
