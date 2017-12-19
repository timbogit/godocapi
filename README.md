# godoc API

## Skills developed

* Building a web API
* Testing a web API
* JSON encoding/decoding
* Exploring godoc
* Extending existing packages or
* Shelling out

## Assignment

Godoc is an amazing tool which can be consumed via the command line or
a web interface. Your goal is to add a web API in front of godoc so its data can be
queried via HTTP/JSON.


# Resources

* http://godoc.org/code.google.com/p/go.tools/cmd/godoc
* http://golang.org/s/oracle-user-manual
* https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo
* https://medium.com/@marcus.olsson/writing-a-go-client-for-your-restful-api-c193a2f4998c
* https://github.com/golang/dep

# To Do:
* introduce concurrency via worker pool for shelling out
  * how to balance load, what to do when high load makes me run out of workers
* caching
  * how to have workers access it.
* package structure
  * where to place models? (unit of reuse for golang, should be ~= to a Java class)
  * see [this article](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)
* make a presentation using [go present](https://godoc.org/golang.org/x/tools/present)
* Apply [functional options](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis) to a `NewCommandService` function
  * That way, the client can move to the top-level
* look at [Akshay's Makefile](https://github.com/akshayjshah/dotfiles/blob/master/Makefile#L8)

# Design
Just some background:
* I am trying to follow the tips in this article regarding package lay-out: https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
* `mux` package is for the web container for an HTTP / RESTful service (... with the idea that, in the future, I could have different protocols / gRPC, thrift, etc.)
* `shell` is for a service (client) implementation that shells out directly (used in the `mux` server, and probably all future servers, as well as in the `cmd/godocapiclient` package)
* `http` is for a service (client) implementation that actually calls the remote HTTP service; this package also has `options.go` (which the `shell` package service doesn't need), which implement the functional options to be passed into `godocapi.NewCommandService` (like HTTP User Agent, port, etc.)
* The idea is that there can also be future different client implementations, as well (for gRPC, etc.), which can reuse the `godocapi.NewCommandService` and pass a separate set of `func(godocapi.CommandService)` in without having to change the implementation of `godocapi.NewCommandService`
* `cmd/godocapi`: contains the code to spin up the (currently HTTP only) service
* `cmd/godocapiclient`: has code to communicate with the HTTP service, as well as via the shell, using a single `godocapi.CommandService` interface typed variable.


