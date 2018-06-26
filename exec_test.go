package exec

import (
	"fmt"
	"os"
	"testing"
)

func TestErrInterpreterNotFound(t *testing.T) {
	envVar := "SHELL"
	err := &ErrInterpreterNotFound{envVar}
	expected := fmt.Sprintf(errMessageInterpreterNotFound, envVar)
	if err.Error() != expected {
		t.Errorf("error returned a wrong message: got %s want %s", err.Error(), expected)
	}
}

func TestCommand(t *testing.T) {
	envVar := "SHELL"
	t.Run(fmt.Sprintf(`calling Command without "%s" environment variable`, envVar), func(t *testing.T) {
		os.Unsetenv(envVar)
		if _, err := Command("foo", "bar"); err == nil {
			t.Errorf(`Command should have thrown an error as "%s" environment variable is not set`, envVar)
		}
	})
	t.Run(fmt.Sprintf(`calling Command with "%s" environment variable`, envVar), func(t *testing.T) {
		os.Setenv(envVar, "bin/sh")
		if _, err := Command("foo", "bar"); err != nil {
			t.Errorf(`Command should not have thrown an error as "%s" environment variable is set`, envVar)
		}
		os.Unsetenv(envVar)
	})
}

func TestCommandContext(t *testing.T) {
	envVar := "SHELL"
	t.Run(fmt.Sprintf(`calling CommandContext without "%s" environment variable`, envVar), func(t *testing.T) {
		os.Unsetenv(envVar)
		if _, err := CommandContext(nil, "foo", "bar"); err == nil {
			t.Errorf(`CommandContext should have thrown an error as "%s" environment variable is not set`, envVar)
		}
	})
	t.Run(fmt.Sprintf(`calling CommandContext with "%s" environment variable`, envVar), func(t *testing.T) {
		os.Setenv(envVar, "bin/sh")
		if _, err := CommandContext(nil, "foo", "bar"); err != nil {
			t.Errorf(`CommandContext should not have thrown an error as "%s" environment variable is set`, envVar)
		}
		os.Unsetenv(envVar)
	})
}
