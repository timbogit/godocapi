package godocapi

type CommandResults struct {
	Command string `json:"command"`
	Output  string `json:"output"`
}


// Client creates a connection to the services.
type Client interface {
	NewCommandService(options ...func(*CommandService)) (*CommandService, error)
}

// CommandService represents a service for managing godoc commands.
type CommandService interface {
	GetSource(arg string) (*CommandResults, error)
}

//func(client *Client) NewCommandService(options ...func(*CommandService)) (*CommandService, error) {
//
//	cmdService := CommandService{}
//
//	for _, option := range options {
//		option(&cmdService)
//	}
//	client.service = &cmdService
//	return &cmdService, nil
//}