package shell

import (
	"github.com/timbogit/godocapi"
)

// Client creates a connection to the services
type CommandClient struct {
	service *CommandService
}


func(client *CommandClient) CommandService() godocapi.CommandService {
	client.service = &CommandService{}

	return client.service
}

