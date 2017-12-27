package mock

import (
	"github.com/timbogit/godocapi"
)


// CommandService represents a mock implementation of godocapi.CommandService.
type CommandService struct {
	CommandFn		func(arg string) (*godocapi.CommandResults, error)
	CommandInvoked	bool
}

func (s *CommandService) GetSource(arg string) (*godocapi.CommandResults, error) {
	s.CommandInvoked = true
	return s.CommandFn(arg)
}