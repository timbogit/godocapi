package main

import (
	"log"
	"github.com/timbogit/godocapi/http"
	"github.com/timbogit/godocapi/shell"
)

func main() {

	log.Println("+++++++++ SHELL CLIENT +++++++++")

	shellClient := shell.CommandClient{}
	shellService := shellClient.CommandService()
	res, err := shellService.GetSource("net/http")

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("\n--- Output ---\n> Command: %s, \n> Source: \n----- Start -----\n%s\n----- End -----\n",
		res.Command, res.Output)


	log.Println("+++++++++ HTTP CLIENT +++++++++")
	httpClient := http.NewCommandClient("localhost:8080", "go-client/1.0.0")
	httpService := httpClient.CommandService()
	res2, err2 := httpService.GetSource("net")

	if err2 != nil {
		log.Fatal(err2.Error())
	}

	log.Printf("\n--- Output ---\n> Command: %s, \n> Source: \n----- Start -----\n%s\n----- End -----\n",
		res2.Command, res2.Output)

}