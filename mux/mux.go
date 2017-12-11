package mux


import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/timbogit/godocapi"

	"github.com/gorilla/mux"
)


type Handler struct {
	CommandService godocapi.CommandService

}

func NewHandler(service godocapi.CommandService) *Handler {
	return &Handler{service}
}

// Display src for 'cmd' argument
func (h *Handler) GetSource(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cmd := params["cmd"]

	log.Printf("requested command is %s", cmd)

	result, err := h.CommandService.GetSource(cmd)

	if err == nil {
		json.NewEncoder(w).Encode(result)
	}

}

func (h *Handler) Serve(host string, r *Router) {
	log.Fatal(http.ListenAndServe(host, r.router))
}

type Router struct {
	router *mux.Router
}


func NewRouter() *Router {
	return &Router{mux.NewRouter()}
}

func (r *Router) HandleGet(slug string, handlerFunc func(http.ResponseWriter, *http.Request)) {
	r.router.HandleFunc(slug, handlerFunc).Methods("GET")
}