MockKit
--------
[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/mockkit)](https://goreportcard.com/report/github.com/gokit/mockkit)

MockKit implements a code generator which automatically generates a go package implementation for a giving declared interface with a mock type.

## Install

```
go get -u github.com/gokit/mockkit
```

## Usage

Running the following commands instantly generates all necessary files and packages for giving code gen.

```go
> mockkit generate
```

## How It works

You annotate any giving interface with `@implement` which marks giving interface has a target for code generation.

Sample below:

```go
// Runner defines an interface for a runner.
// @implement
type Runner struct {
	NextRun() (float64, error)
}
```
