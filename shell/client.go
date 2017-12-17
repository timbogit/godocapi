package shell

// Client creates a connection to the services
type CommandClient struct {
	service *CommandService
}

func(client *CommandClient) NewCommandService(options ...func(*CommandService)) (*CommandService, error) {

	cmdService := CommandService{}

	for _, option := range options {
		option(&cmdService)
	}
	client.service = &cmdService
	return &cmdService, nil
}