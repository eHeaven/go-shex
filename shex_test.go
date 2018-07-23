package shex

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestCommand(t *testing.T) {
	envVar := "SHELL"
	t.Run(fmt.Sprintf(`calling Command without "%s" environment variable`, envVar), func(t *testing.T) {
		os.Unsetenv(envVar)
		if _, err := Command("foo", "bar"); err == nil {
			t.Errorf(`Command should have thrown an error as "%s" environment variable is not set`, envVar)
		}
	})
	t.Run(fmt.Sprintf(`calling Command with "%s" environment variable`, envVar), func(t *testing.T) {
		os.Setenv(envVar, "/bin/sh")
		if _, err := Command("foo", "bar"); err != nil {
			t.Errorf(`Command should not have thrown an error as "%s" environment variable is set: got "%s"`, envVar, err.Error())
		}
		os.Unsetenv(envVar)
	})
}

func TestCommandContext(t *testing.T) {
	envVar := "SHELL"
	t.Run(fmt.Sprintf(`calling CommandContext without "%s" environment variable`, envVar), func(t *testing.T) {
		os.Unsetenv(envVar)
		if _, err := CommandContext(context.TODO(), "foo", "bar"); err == nil {
			t.Errorf(`CommandContext should have thrown an error as "%s" environment variable is not set`, envVar)
		}
	})
	t.Run(fmt.Sprintf(`calling CommandContext with "%s" environment variable`, envVar), func(t *testing.T) {
		os.Setenv(envVar, "/bin/sh")
		if _, err := CommandContext(context.TODO(), "foo", "bar"); err != nil {
			t.Errorf(`CommandContext should not have thrown an error as "%s" environment variable is set: got "%s"`, envVar, err.Error())
		}
		os.Unsetenv(envVar)
	})
}

func TestSafeCommand(t *testing.T) {
	t.Run("calling SafeCommand", func(t *testing.T) {
		if _, err := SafeCommand("foo", "bar"); err != nil {
			t.Errorf(`SafeCommand should not have thrown an error: got "%s"`, err.Error())
		}
	})
}

func TestSafeCommandContext(t *testing.T) {
	t.Run("calling SafeCommandContext", func(t *testing.T) {
		if _, err := SafeCommandContext(context.TODO(), "foo", "bar"); err != nil {
			t.Errorf(`SafeCommandContext should not have thrown an error: got "%s"`, err.Error())
		}
	})
}

func TestRun(t *testing.T) {
	t.Run("testing running an external command with Command", func(t *testing.T) {
		envVar := "SHELL"
		os.Setenv(envVar, "/bin/sh")
		cmd, err := Command("echo", "Hello world")
		if err != nil {
			t.Fatalf(`An unexpected error occurred while creating an instance of exec.Cmd: got "%s"`, err.Error())
		}
		if err := cmd.Run(); err != nil {
			t.Errorf(`Run should not have thrown an error: got "%s"`, err.Error())
		}
		os.Unsetenv(envVar)
	})
	t.Run("testing running an external command with SafeCommand", func(t *testing.T) {
		cmd, err := SafeCommand("echo", "Hello world")
		if err != nil {
			t.Fatalf(`An unexpected error occurred while creating an instance of exec.Cmd: got "%s"`, err.Error())
		}
		if err := cmd.Run(); err != nil {
			t.Errorf(`Run should not have thrown an error: got "%s"`, err.Error())
		}
	})
}
