package godocapi

type CommandResults struct {
	Command string `json:"command"`
	Output  string `json:"output"`
}


// Client creates a connection to the services.
type Client interface {
	CommandService() CommandService
}

// CommandService represents a service for managing godoc commands.
type CommandService interface {
	GetSource(arg string) (*CommandResults, error)
}