# Go Bootstrap Kata

This is kata to get you taste of Go with some hands on tasks.

If anytime you are stuck you can check the [docs](https://pkg.go.dev/), and specifically the [standard library docs](https://pkg.go.dev/std) since we will try to not add other dependencies.
And in case you cannot tell from where an identifier comes from it might be a [builtin](https://pkg.go.dev/builtin).

Anytime you have a doubt on a concept, be it syntax or semantic or both, the [Effective Go](https://go.dev/doc/effective_go) document will probably help, keep it handy.

## Task 0 - Getting Ready

If you did not set up a Go development environment this is your first task 😉.
Installing instruction for the Go compiler and tools are available on the [official Go website](https://go.dev/doc/install)

Go may also be available as a package for your OS, e.g. on MacOS is available via `brew isntall go`.

Remember to check the installation was successful by typing `go version` in your terminal.


Go has an official [Language Server](https://github.com/golang/tools/tree/master/gopls#gopls-the-go-language-server), `gopls`, with plugins available for popular editors.


## Task 1 - Hello World

Let's compile and run a Hello World program and get familiar with the go tool.
The code for this task is in `cmd/hello/`.

Here is a list of things you may want to try:
- open `main.go` in your editor and have a look around
- run `go run .`
- run `go run main.go`
- run `go build` to compile the `hello` command, check the directory afterward and run your new executable
- run `go env` to see all the go environment variables
- run `go help` to see more go sub-commands, there are a few so you might want to use a pager


## Task 2 - Hello Tests

Now that we have some basic familiarity with the `go` tool it is time for some tests!

The `users/users.go` file defines a go package and some tests for the type defined in the package can be found in `users/users_test.go`.
Go provides a [testing package in the standard library](https://pkg.go.dev/testing) and it is integrated with the `go` tool making it a quite comprehensive testing framework.

The canonical way of testing go packages is to run `go test` in a directory containing a package, or run `go test ./...` in the root directory of a module to run all the test in the module.

Is there something wrong with the `users` package? Let's find out!

- from the root directory of this repository try to run `go test ./...` to test the whole module
- cd into the `users` directory and run `go test` to test only the `users` package
- run `go help test`
- run `go test -json`


Test coverage is also included (as are benchmarking and fuzzing) and works out of the box, try some of these commands

- run `go help testflag`
- run `go test -coverprofile=coverage.out`
- run `go tool cover -func=coverage.out` 
- run `go tool cover -html=coverage.out` 
- see the whole cover story at [go.dev/blog/cover](https://go.dev/blog/cover)




## Task 3 - Hello Web

In this task we are going to play around with a small web server, the web server is not accomplishing much yet but it will be useful to show some nice Go features.

- cd into `cmd/hello-web`
- you should know by now how to compile and run  the code ;)
- have a look at the `better-logs.go` file, can you improve the log middleware as the comments mentioned?

## Task 4 - Hello JSON

WIP: tyny JSON API  maybe just one endpoint /items/{ID} to add an item or list all the items, write some tests for it.
Goals:
- use JSON
- see how to test HTTP endpoints

## Task 5 - Hello Goroutines

WIP: expand on the JSON API maybe add an enpoint to process the items /items/process/ that accepts a list of items and process them in a goroutine.
Goals:
- use pool of workers

## Task 6 - Hello GOOS and GOARCH

WIP: see how to cross compile, maybe also the resouce embedding 

## Task 7

Beer

