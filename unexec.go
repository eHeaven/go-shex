/*
Package unexec is a simple library easing the use of os/exec package for running cross-platform external commands.

While using the os/exec package, you may have encountered some consistency issues:
a command which was working fine on your command line interpreter fails miserably while calling it
with the said package.

To address this common problem, the go-unexec library tries to detect your default command line
interpreter by looking for the SHELL environment variable on UNIX systems or COMSPEC environment variable
on Windows.

So previously your code might have looked like this:

 import "os/exec"

 func main() {
	 cmd := exec.Command("echo", "Hello world")
	 // will run "echo Hello world".
 }

With this package:

 import unexec "github.com/thegomachine/go-unexec"

 func main() {
	 cmd, err := unexec.Command("echo", "Hello world")
	 // will run "/bin/sh -c echo Hello world" (or "/bin/zsh -c echo Hello world" etc.)
	 // on UNIX systems or "cmd.exe /c echo Hello world" on Windows.
 }
*/
package unexec

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Command is a wrapper function around exec.Command function from os/exec package.
// If the command line interpreter is not found, throws an ErrInterpreterNotFound error.
func Command(name string, arg ...string) (*exec.Cmd, error) {
	return makeCommand(nil, name, arg...)
}

// CommandContext is a wrapper function around exec.CommandContext function from os/exec package.
// If the command line interpreter is not found, throws an ErrInterpreterNotFound error.
func CommandContext(ctx context.Context, name string, arg ...string) (*exec.Cmd, error) {
	return makeCommand(ctx, name, arg...)
}

func makeCommand(ctx context.Context, name string, arg ...string) (*exec.Cmd, error) {
	interpreter, err := lookForInterpreter()
	if err != nil {
		return nil, err
	}

	var args []string
	args = append(args, name)
	args = append(args, arg...)

	if ctx != nil {
		return exec.CommandContext(ctx, interpreter[0], interpreter[1], strings.Join(args, " ")), nil
	}

	return exec.Command(interpreter[0], interpreter[1], strings.Join(args, " ")), nil
}

// ErrInterpreterNotFound is thrown when the command line interpreter was not found.
type ErrInterpreterNotFound struct {
	envVar string
}

const errMessageInterpreterNotFound = `"%s" is a required environment variable: it allows to know which command line interpreter to use for running a command`

func (e *ErrInterpreterNotFound) Error() string {
	return fmt.Sprintf(errMessageInterpreterNotFound, e.envVar)
}

func lookForInterpreter() ([]string, error) {
	var (
		envVar string
		flag   string
	)

	if runtime.GOOS == "windows" {
		envVar = "COMSPEC"
		flag = "/c"
	} else {
		envVar = "SHELL"
		flag = "-c"
	}

	binary := os.Getenv(envVar)
	if binary == "" {
		return nil, &ErrInterpreterNotFound{envVar}
	}

	return []string{binary, flag}, nil
}
