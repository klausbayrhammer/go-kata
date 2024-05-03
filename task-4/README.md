## Task 4 - Hello Web

In a Go source tree the canonical location of code for executables is under the `cmd` directory.
And each executable, or command, in in its own subdirectory:

```
awesome_go_codebase
├── cmd
│   ├── command1
│   │   └── main.go
│   └── command2
│       └── main.go
├── package1
│   └── package1_code.go
├── package2
│   └── package2_code.go
├── go.mod
├── go.sum
└── README.md
```

In this task we are going to play around with a web server, meaning we are going to compile and run an executable.
In `task-1` we violated this rule, but it was for a good cause!

In order to follow the conventions this time all the code is in  the `cmd/taks-4` directory.

