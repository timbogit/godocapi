package http

import (
	"net/http"
	"net/url"
)

// Client creates a connection to the services
type CommandClient struct {
	service *CommandService
}

func(client *CommandClient) SetBaseUrl(host string) func(*CommandService) {
	return func(service *CommandService) {
		service.BaseURL = &url.URL{Scheme: "http", Host: host}
	}
}

func(client *CommandClient) SetAgent(agent string) func(*CommandService) {
	return func(service *CommandService) {
		service.UserAgent = agent
	}
}

func(client *CommandClient) SetHttpClient(c *http.Client) func(*CommandService) {
	return func(service *CommandService) {
		if c != nil {
			service.httpClient = c
		} else {
			service.httpClient = http.DefaultClient
		}
	}
}

func(client *CommandClient) NewCommandService(options ...func(*CommandService)) (*CommandService, error) {

	client.service = &CommandService{}

	for _, option := range options {
		option(client.service)
	}
	return client.service, nil
}