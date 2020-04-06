# Go Find Up

[![GoDoc](https://godoc.org/github.com/ojizero/gofindup?status.svg)](https://godoc.org/github.com/ojizero/gofindup)
![Run tests](https://github.com/ojizero/gofindup/workflows/Run%20tests/badge.svg?branch=master)

> Find a file or directory by walking up parent directories.

## Usage

Install latest version

```sh
go get github.com/ojizero/gofindup
```

Import it into your code

```go
import "github.com/ojizero/gofindup"
```

This package only exposes 2 functions, `Findup` and `FindupFrom`,

```go
// looks the file recursively in parents starting from "./"
gofindup.Findup("some-file-or-directory")

// looks the file recursively in parents starting from "./some-starting-directory"
gofindup.FindupFrom("some-file-or-directory", "some-starting-directory")
```
