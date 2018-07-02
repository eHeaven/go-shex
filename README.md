<p align="center">
    <img src="https://user-images.githubusercontent.com/8983173/41920404-f4baf4e2-7960-11e8-8880-6b54bcef12e2.png" alt="Logo" width="200" height="200" />
</p>
<h3 align="center">go-shex</h3>
<p align="center">A simple package for creating <a href="https://golang.org/pkg/os/exec/#Cmd">exec.Cmd</a> instances which use the current command interpreter (the shell)</p>
<p align="center">
    <a href="https://travis-ci.org/thegomachine/go-shex">
        <img src="https://travis-ci.org/thegomachine/go-shex.svg?branch=master" alt="Travis CI">
    </a>
    <a href="https://godoc.org/github.com/thegomachine/go-shex">
        <img src="https://godoc.org/github.com/thegomachine/go-shex?status.svg" alt="GoDoc">
    </a>
    <a href="https://goreportcard.com/report/thegomachine/go-shex">
        <img src="https://goreportcard.com/badge/github.com/thegomachine/go-shex" alt="Go Report Card">
    </a>
    <a href="https://codecov.io/gh/thegomachine/go-shex/branch/master">
        <img src="https://codecov.io/gh/thegomachine/go-shex/branch/master/graph/badge.svg" alt="Codecov">
    </a>
</p>

---

While using the `os/exec` package, you may have encountered some consistency issues:
a command which was working fine on your command interpreter fails miserably while calling it
with the said package.

To address this common problem, the go-shex package tries to detect your default command
interpreter by looking for the `SHELL` environment variable on UNIX systems or `COMSPEC` environment variable
on Windows.

## Installation

```bash
$ go get github.com/thegomachine/go-shex
```

## Usage

So previously your code might have looked like this:

```golang
import "os/exec"

func main() {
    cmd := exec.Command("echo", "Hello world")
    // will run "echo Hello world".
 }
```

With this package:

```golang
import shex "github.com/thegomachine/go-shex"

func main() {
    cmd, err := shex.Command("echo", "Hello world")
    // will run "/bin/sh -c echo Hello world" (or "/bin/zsh -c echo Hello world" etc.)
    // on UNIX systems or "cmd.exe /c echo Hello world" on Windows.
}
```

See [GoDoc](https://godoc.org/github.com/thegomachine/go-shex) for full documentation.