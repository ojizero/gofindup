# Go Find Up

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
gofindup.Findup("some-file-or-directory", "some-starting-directory")
```
