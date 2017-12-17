package main

import (
	"log"
	"github.com/timbogit/godocapi/http"
	"github.com/timbogit/godocapi/shell"
)

func main() {

	log.Println("+++++++++ SHELL CLIENT +++++++++")

	shellClient := &shell.CommandClient{}
	shellService, _ := shellClient.NewCommandService()
	res, err := shellService.GetSource("net/http")

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("\n--- Output ---\n> Command: %s, \n> Source: \n----- Start -----\n%s\n----- End -----\n",
		res.Command, res.Output)


	log.Println("+++++++++ HTTP CLIENT +++++++++")
	httpClient := http.CommandClient{}
	httpService, _ := httpClient.NewCommandService(
		httpClient.SetAgent("go-client/1.0"),
		httpClient.SetBaseUrl("localhost:8080"),
		httpClient.SetHttpClient(nil))
	res2, err2 := httpService.GetSource("net")

	if err2 != nil {
		log.Fatal(err2.Error())
	}

	log.Printf("\n--- Output ---\n> Command: %s, \n> Source: \n----- Start -----\n%s\n----- End -----\n",
		res2.Command, res2.Output)

}