package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/timbogit/godocapi"
)


// CommandService represents a service for managing commands.
type CommandService struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

func (s *CommandService) GetSource(arg string) (*godocapi.CommandResults, error) {
	rel := &url.URL{Path: fmt.Sprintf("/godoc/src/%s", arg)}
	u := s.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.UserAgent)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var sourceCommandResult godocapi.CommandResults

	err = json.NewDecoder(resp.Body).Decode(&sourceCommandResult)
	if err == nil {
		return &sourceCommandResult, nil
	}

	return nil, err
}
