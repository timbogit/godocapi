package godocapi

type CommandResults struct {
	Command string `json:"command"`
	Output  string `json:"output"`
}

// CommandService represents a service for managing godoc commands.
type CommandService interface {
	GetSource(arg string) (*CommandResults, error)
}

func NewCommandService(initial CommandService, options ...func(CommandService)) (CommandService, error) {
	for _, option := range options {
		option(initial)
	}
	return initial, nil
}