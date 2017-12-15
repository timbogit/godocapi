package main

import (
	"log"
	"github.com/timbogit/godocapi/shell"
)

func main() {
	client := shell.CommandClient{}
	srv := client.CommandService()
	res, err := srv.GetSource("net/http")

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("\n--- Output ---\n> Command: %s, \n> Source: \n----- Start -----\n%s\n----- End -----\n",
		res.Command, res.Output)
}