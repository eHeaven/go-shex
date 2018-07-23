/*
Package shex is simple package for creating https://golang.org/pkg/os/exec/#Cmd instances which use the current command interpreter (the shell).

While using the os/exec package, you may have encountered some consistency issues:
a command which was working fine on your command interpreter fails miserably while calling it
with the said package.

To address this common problem, the go-shex library tries to detect your default command
interpreter by looking for the SHELL environment variable on UNIX systems or COMSPEC environment variable
on Windows.

So previously your code might have looked like this:

 import "os/exec"

 func main() {
	 cmd := exec.Command("echo", "Hello world")
	 // will run "echo Hello world".
 }

With this package:

 import shex "github.com/thegomachine/go-shex"

 func main() {
	 cmd, err := shex.Command("echo", "Hello world")
	 // will run "/bin/sh -c echo Hello world" (or "/bin/zsh -c echo Hello world" etc.)
	 // on UNIX systems or "cmd.exe /c echo Hello world" on Windows.

	 // if you don't want auto-detection, you may also use:
	 cmd, err := shex.SafeCommand("echo", "Hello world")
	 // will run "/bin/sh -c echo Hello world" on UNIX systems
	 // or "cmd.exe /c echo Hello world" on Windows.
 }
*/
package shex

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	auto = iota
	safe
)

const (
	defaultEnvVar = "SHELL"
	defaultShell  = "/bin/sh"
	defaultFlag   = "-c"
)

// Command is a wrapper function around exec.Command function from os/exec package.
// If the command line interpreter is not found, throws an error.
func Command(name string, arg ...string) (*exec.Cmd, error) {
	return makeCommand(nil, auto, name, arg...)
}

// CommandContext is a wrapper function around exec.CommandContext function from os/exec package.
// If the command line interpreter is not found, throws an error.
func CommandContext(ctx context.Context, name string, arg ...string) (*exec.Cmd, error) {
	return makeCommand(ctx, auto, name, arg...)
}

// SafeCommand is a wrapper function around exec.Command function from os/exec package.
// Unlike Command function, it will always use "/bin/sh" on UNIX systems or "cmd.exe" on Windows.
// If the command line interpreter is not found, throws an error.
func SafeCommand(name string, arg ...string) (*exec.Cmd, error) {
	return makeCommand(nil, safe, name, arg...)
}

// SafeCommandContext is a wrapper function around exec.CommandContext function from os/exec package.
// Unlike CommandContext function, it will always use "/bin/sh" on UNIX systems or "cmd.exe" on Windows.
// If the command line interpreter is not found, throws an error.
func SafeCommandContext(ctx context.Context, name string, arg ...string) (*exec.Cmd, error) {
	return makeCommand(ctx, safe, name, arg...)
}

func makeCommand(ctx context.Context, mode int, name string, arg ...string) (*exec.Cmd, error) {
	var args []string
	args = append(args, name)
	args = append(args, arg...)
	shell, err := fetchShell(mode)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		return exec.CommandContext(ctx, shell, defaultFlag, strings.Join(args, " ")), nil
	}
	return exec.Command(shell, defaultFlag, strings.Join(args, " ")), nil
}

func fetchShell(mode int) (string, error) {
	shell := defaultShell
	if mode == auto {
		shell := os.Getenv(defaultEnvVar)
		if shell == "" {
			return "", fmt.Errorf(`"%s" is a required environment variable: it allows to know which command interpreter to use for running the external command`, defaultEnvVar)
		}
	}
	if _, err := exec.LookPath(shell); err != nil {
		return "", err
	}
	return shell, nil
}
