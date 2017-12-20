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
	// Keep these unexported! Part of the goal of functional options is to let
	// users configure these details without exporting the struct fields.
	baseURL   *url.URL
	userAgent string

	httpClient *http.Client
}

func New(opts ...Option) *CommandService {
	s := &CommandService{
		baseURL:    &url.URL{Scheme: "http", Host: ":80"}, // reasonable default
		httpClient: http.DefaultClient,
	}
	for _, o := range opts {
		o.apply(s)
	}
	return s
}

func (s *CommandService) GetSource(arg string) (*godocapi.CommandResults, error) {
	rel := &url.URL{Path: fmt.Sprintf("/godoc/src/%s", arg)}
	u := s.baseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.userAgent)

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
