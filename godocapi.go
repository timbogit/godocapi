package godocapi

type CommandResults struct {
	Command string `json:"command"`
	Output  string `json:"output"`
}


//// Client creates a connection to the services.
//type Client interface {
//	NewCommandService(options ...func(*CommandService)) (*CommandService, error)
//}

// CommandService represents a service for managing godoc commands.
type CommandService interface {
	GetSource(arg string) (*CommandResults, error)
}

// Client creates a connection to the services.
type CommandClient struct {
	Service *CommandService
}

func NewCommandService(initial CommandService, options ...func(CommandService)) (CommandService, error) {
	for _, option := range options {
		option(initial)
	}
	return initial, nil
}