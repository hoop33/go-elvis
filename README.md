# Elvis

> Elvis and ternary operators for go

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](http://opensource.org/licenses/MIT)
[![Issues](https://img.shields.io/github/issues/hoop33/go-elvis.svg)](https://github.com/hoop33/go-elvis/issues)
[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/hoop33/go-elvis)
[![Go Report Card](https://goreportcard.com/badge/github.com/hoop33/go-elvis)](https://goreportcard.com/report/github.com/hoop33/go-elvis)
## Introduction

Go doesn't have the [Elvis operator](https://en.wikipedia.org/wiki/Elvis_operator) or the [Ternary operator](https://en.wikipedia.org/wiki/%3F:). This package provides both operators as functions. It's a what-if project that may turn out to be useful.

## Installation

```sh
$ go get -u github.com/hoop33/go-elvis
```

## Usage

See the `examples` directory for example usages

### Elvis

The `Elvis` function takes two arguments and returns an empty interface. If the first argument is `nil` or the zero value, it returns the second argument. Otherwise, it returns the first argument. Note that you probably want to convert the empty interface to the type you're passing. Examples:

```go
var a string
a = "a"
a = elvis.Elvis(a, "b").(string) // a is now "a"

a = ""
a = elvis.Elvis(a, "b").(string) // a is now "b"

var n int
n = 1
n = elvis.Elvis(n, 2).(int) // n is now 1

n = 0
n = elvis.Elvis(n, 2).(int) // n is now 2

var p1, p2, p3 *string
p := "Hello"
p1 = nil
p2 = &p
p3 = elvis.Elvis(p1, p2).(*string) // p3 is now p2
```

### Ternary

The `Ternary` function takes three arguments: a conditional (`bool`) and two other arguments and returns an empty interface. If the conditional is `true`, it returns the first argument. Otherwise, it returns the second argument. Note that you probably want to convert the empty interface to the type you're passing. Examples:

```go
x := 33
y := elvis.Ternary(x == 33, "Hello", "Goodbye").(string) // y is now "Hello"
y = elvis.Ternary(x == 11, "Hello", "Goodbye").(string) // y is now "Goodbye"
```

## Contributing

Please note that this project is released with a [Contributor Code of Conduct](http://contributor-covenant.org/). By participating in this project you agree to abide by its terms. See [CODE_OF_CONDUCT](CODE_OF_CONDUCT.md) file.

Contributions are welcome! Please open pull requests with code that passes all the checks. See *Building* for more information.

### Building

You must have a working Go development environment to contribute code.

The included `Makefile` performs various checks on the code. To get started, run:

```sh
$ make deps
```

This will install the necessary dependencies. You should have to do this only once.

Then, you can run:

```sh
$ make
```

To run the code checks and tests.

## License

Copyright &copy; 2017 Rob Warner

Licensed under the [MIT License](https://hoop33.mit-license.org/)
