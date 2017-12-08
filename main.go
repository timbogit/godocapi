package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

type Source struct {
	Code        string   `json:"code,omitempty"`
}

// Display src for 'cmd' argument
func GetSource(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cmd := params["cmd"]

	output, err := exec.Command("godoc", "-src", cmd).Output()

	if err == nil {
		source := Source{Code: string(output)}
		json.NewEncoder(w).Encode(source)
	}

}

// our main function
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/godoc/src/{cmd}", GetSource).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}