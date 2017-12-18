package http

import (
	"net/http"
	"net/url"

	"github.com/timbogit/godocapi"
)


func SetBaseUrl(host string) func(godocapi.CommandService) {
	return func(service godocapi.CommandService) {
		service.(*CommandService).BaseURL = &url.URL{Scheme: "http", Host: host}
	}
}

func SetAgent(agent string) func(godocapi.CommandService) {
	return func(service godocapi.CommandService) {
		service.(*CommandService).UserAgent = agent
	}
}

func SetHttpClient(c *http.Client) func(godocapi.CommandService) {
	return func(service godocapi.CommandService) {
		if c != nil {
			service.(*CommandService).httpClient = c
		} else {
			service.(*CommandService).httpClient = http.DefaultClient
		}
	}
}
