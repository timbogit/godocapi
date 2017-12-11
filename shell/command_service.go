package shell


import (
	"github.com/timbogit/godocapi"
	"os/exec"
)

// CommandService represents a service for managing commands.
type CommandService struct {
}

func (s *CommandService) GetSource(arg string) (*godocapi.CommandResults, error) {
	output, err := exec.Command("godoc", "-src", arg).Output()

	var sourceCommandResult godocapi.CommandResults
	if err == nil {
		sourceCommandResult = godocapi.CommandResults{Command: arg, Output:string(output)}
		return &sourceCommandResult, nil
	}

	return nil, err
}