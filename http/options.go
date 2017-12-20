package http

import (
	"net/http"
	"net/url"
)

type Option interface {
	// Keeping this unexported makes it completely opaque to the user, which
	// means that you can change it later.
	apply(*CommandService)
}

// This is an idiomatic bit of Go trickery - functions can have methods! The
// standard library's http.HandlerFunc is the most common example. Now we have
// a small adapter that makes it easy for us to turn anonymous functions into
// Options.
type optionFunc func(*CommandService)

func (f optionFunc) apply(s *CommandService) { f(s) }

func SetBaseUrl(host string) Option {
	return optionFunc(func(service *CommandService) {
		service.baseURL = &url.URL{Scheme: "http", Host: host}
	})
}

func SetAgent(agent string) Option {
	return optionFunc(func(service *CommandService) {
		service.userAgent = agent
	})
}

func SetHttpClient(c *http.Client) Option {
	return optionFunc(func(service *CommandService) {
		// Options typically don't contain the logic to fall back to the default
		// HTTP client if the user passes a nil - part of the point of this
		// pattern is to avoid explicitly passing nil to trigger the default
		// behavior. Instead, move that logic into the constructor.
		//
		// If the user explicitly passes a nil here, it's perfectly fine to panic
		// later on with a nil pointer dereference - idiomatic Go code doesn't
		// contort itself to protect users from themselves.
		service.httpClient = http.DefaultClient
	})
}
