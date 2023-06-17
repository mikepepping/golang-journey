# Setup Project

## To setup a new project

You need to create a new go module.
This will be your entire project.

The module name should ideally be where your code will sit.
So for this project it would be

```bash
go mod init github.com/mikepepping/golang-journey/setup-project
```

## Building

To build simple run `go build <<entry file>>`

For this project:
```bash
go build main.go
```

## Running

To build and then run a go project simple run `go run <<entry file>>`

For example
```bash
go run main.go
```

