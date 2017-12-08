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
* https://github.com/golang/dep

# To Do:
* introduce concurrency via worker pool for shelling out
  * how to balance load, what to do when high load makes me run out of workers
* caching
  * how to have workers access it.
* package structure
  * where to place models? (unit of reuse for golang, should be ~= to a Java class)
  * see [this article](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)
* make a presentation using (go present][https://godoc.org/golang.org/x/tools/present]
* look at [Akshay's Makefile](https://github.com/akshayjshah/dotfiles/blob/master/Makefile#L8)

