# Go Bootstrap Kata

This is kata to get you a taste of Go with some hands on tasks.

If anytime you are stuck you can check the [docs](https://pkg.go.dev/), and specifically the [standard library docs](https://pkg.go.dev/std) since we will try to not add other dependencies.
And in case you cannot tell from where an identifier comes from it might be a [builtin](https://pkg.go.dev/builtin).

Anytime you have a doubt on a concept, be it syntax or semantic or both, the [Effective Go](https://go.dev/doc/effective_go) document will probably help, keep it handy.


There are several tasks for you, none of them is very specific just let the comments in the code guide you and have fun!

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



## Task 2 - Hello tests

Now that we have some basic familiarity with the `go` tool it is time for some tests!

Go provides a [testing package in the standard library](https://pkg.go.dev/testing) and it is integrated with the `go` tool making it a quite comprehensive testing framework.

The canonical way of testing go packages is to run `go test` in a directory containing a package, or run `go test ./...` in the root directory of a module to run all the test in the module.

Is there something wrong with the `address` package in the `task-2` directory? 
Can you improve the code and write some more test?
Can you check the test coverage?

Some tips:

- from the root directory of this repository try to run `go test ./...` to test the whole module
- from the root directory of this repository try to run `go test ./task-2` to test only the package in that directory
- cd into the `task-2` directory and run `go test`
- run `go help test`
- run `go test -json`


Test coverage is also included (as are benchmarking and fuzzing) and works out of the box, try some of these commands

- run `go help testflag`
- run `go test -coverprofile=coverage.out`
- run `go tool cover -func=coverage.out` 
- run `go tool cover -html=coverage.out` 
- see the whole cover story at [go.dev/blog/cover](https://go.dev/blog/cover)


## Task 3 - Hello Composition

This task will make you familiarize with how Go avoid inheritance in favor of composition. 
Another interesting feature of Go is the structural sub-typing that interfaces provide, along with the Liskov substitution.


With your help we want to build an online shop to make Jeff green with envy!
I this task we are going to use the `Address` from `task-2` and provide our `User`s with a billing and shipping address.

Also we need our users to be able to login, maybe we can use and interface to delegate the actual logic to different specific implementations.
 
Some tips:

Given a `Count` type:

```go
type Count struct{
   Value int 
}


func(c *Count) Inc(){
    c.Value +=1
}

```

We can embed it in another type:

```go 
type Foo struct {
    Bar string
    Count
}
```

The embedding type will have the methods of the embedded:

```go
f := Foo{
    Bar: "Bar",
    // Count is created with its zero-value
}

f.Inc() // now f.Count.Value == 1 
fmt.Println(f.Value) // we can skip to the inner values of the embedded type
```

Embedding can be done with Interfaces as well, when we create a type embedding an interface we need to provide a type implementing that interface.

```go
type Baz struct {
    Qux string
    io.Writer // this is an embedded interface
}

b := Baz{
    Qux: "Qux",
    Writer: bytes.NewBuffer([]byte{'H','e','l','l','o'}), // bytes.Buffer implements io.Writer
}

```

## Task 4 - Hello Web

In this task we are going to play around with a small web server, the web server is not accomplishing much yet but it will be useful to show some nice Go features.

- cd into `cmd/hello-web`
- you should know by now how to compile and run  the code ;)
- have a look at the `better-logs.go` file, can you improve the log middleware as the comments mentioned?


## Task 5 - Hello Goroutines

**David Attenborough voice**
Ah the fabled Goroutines... not a system thread and not a coroutine, more like a flock of concurrent subroutines multiplexed on a pool of system threads...

Go has a concurrency model not based on `async` `await` but based on CSP.
The [CSP model](https://www.cs.cmu.edu/~crary/819-f09/Hoare78.pdf)(Commicating Sequential Processes) was devised by Tony Hoare ~50 years ago.
The same [Tony Hoare](https://en.wikipedia.org/wiki/Tony_Hoare) is the one that admitted to the "billion dollar mistake" by introducing `NULL` pointers.
The irony is not wasted on me for that Go has null pointers...

Back to the Goroutines 
To use the fa goroutines, select, and channels.
- use pool of workers

## Task 6 - Hello GOOS and GOARCH

Cross compile in Go is as simple as setting `GOOS` and `GOARCH` and running `go build`.

To try it out is better to see the available values in you Go installation.
Run `go tool dist list` to see them all.

Try to build a binary for another OS or CPU architecture!


## Task 7 - Beer

It's Beer o'clock!



### References 

- [Phil Wadler: Featherweight Go](https://www.youtube.com/watch?v=Dq0WFigax_c)
