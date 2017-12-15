package http

import (
	"net/http"
	"net/url"
	"github.com/timbogit/godocapi"
)

// Client creates a connection to the services
type CommandClient struct {
	service *CommandService
}


func(client *CommandClient) CommandService() godocapi.CommandService {
	return client.service
}

func NewCommandClient(host string, agent string) *CommandClient {
	return &CommandClient{
		&CommandService{
			&url.URL{Scheme: "http", Host: host}, agent, http.DefaultClient }}
}