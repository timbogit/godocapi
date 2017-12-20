package godocapi

type CommandResults struct {
	Command string `json:"command"`
	Output  string `json:"output"`
}

// CommandService represents a service for managing godoc commands.
type CommandService interface {
	GetSource(arg string) (*CommandResults, error)
}
