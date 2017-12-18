package main

import (
	"log"

	"github.com/timbogit/godocapi"
	"github.com/timbogit/godocapi/http"
	"github.com/timbogit/godocapi/shell"
)

func main() {

	log.Println("+++++++++ SHELL CLIENT +++++++++")
	service, _ := godocapi.NewCommandService(&shell.CommandService{})

	res, err := service.GetSource("net/http")

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("\n--- Output ---\n> Command: %s, \n> Source: \n----- Start -----\n%s\n----- Shell End -----\n",
		res.Command, res.Output)


	log.Println("+++++++++ HTTP CLIENT +++++++++")
	service, _ = godocapi.NewCommandService(&http.CommandService{},
		http.SetAgent("go-client/1.0"),
		http.SetBaseUrl("localhost:8080"),
		http.SetHttpClient(nil) )
	res2, err2 := service.GetSource("net")

	if err2 != nil {
		log.Fatal(err2.Error())
	}

	log.Printf("\n--- Output ---\n> Command: %s, \n> Source: \n----- Start -----\n%s\n----- HTTP End -----\n",
		res2.Command, res2.Output)

}