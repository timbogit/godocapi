package main

import (
"github.com/timbogit/godocapi/mux"
"github.com/timbogit/godocapi/shell"
)

func main() {
	cmdService := shell.CommandService{}
	handler := mux.NewHandler(&cmdService)
	router := mux.NewRouter()
	router.HandleGet("/godoc/src/{cmd}", handler.GetSource)

	handler.Serve(":8080", router)

}