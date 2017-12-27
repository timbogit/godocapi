package mux_test

import (
	"net/http/httptest"
	"net/http"
	"testing"

	"github.com/timbogit/godocapi"
	"github.com/timbogit/godocapi/mock"
	"github.com/timbogit/godocapi/mux"
)

func executeRequest(r *mux.Router, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r.Serve(w, req)

	return w
}

func TestHandler(t *testing.T) {
	// Injecting mock user Service into the handler
	var commandService mock.CommandService
	handler := mux.NewHandler(&commandService)
	router := mux.NewRouter()
	router.HandleGet("/test/src/{cmd}", handler.GetSource)

	// Mock out GetSource() call
	commandService.CommandFn = func(arg string) (*godocapi.CommandResults, error) {
		if arg != "test_command" {
			t.Fatalf("unexpected arg: %s", arg)
		}

		return &godocapi.CommandResults{Command: "test_command", Output: "this is a test!"}, nil
	}

	r, _ := http.NewRequest("GET", "/test/src/test_command", nil)
	executeRequest(router, r)

	// Validate mock.
	if !commandService.CommandInvoked {
		t.Fatal("expected GetSource() to be invoked")
	}
}
