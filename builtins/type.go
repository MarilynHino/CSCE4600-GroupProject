package builtins

import (
	"errors"
	"fmt"
	"os/exec"
)

// Define an error variable for an invalid number of arguments.
var (
	ErrTypeInvalidArgCount = errors.New("invalid argument count")
)

// To checks whether a command is a shell built-in command.
func TypeCommand(args ...string) (string, error) {
	// Ensure that only one argument was passed to the function.
	if len(args) != 1 {
		return "", fmt.Errorf("%w: expected one argument (command name)", ErrTypeInvalidArgCount)
	}

	cmd := args[0]

	// Check if the command is a shell built-in.
	if IsBuiltin(cmd) {
		return fmt.Sprintf("%s is a shell builtin", cmd), nil
	}

	// Using the exec.LookPath function to find the path to the command.
	path, err := exec.LookPath(cmd)
	if err != nil {
		// If the command was not found, return an error.
		return "", fmt.Errorf("%s: command not found", cmd)
	}

	// Return a message indicating that the command was found.
	return fmt.Sprintf("%s is %s", cmd, path), nil
}

// Checks whether a command is a shell built-in command.
func IsBuiltin(cmd string) bool {
	// Define a list of known shell built-ins.
	builtins := []string{"cd", "type", "shift", "export", "exit", "env"}
	// Iterate over the list of built-ins.
	for _, b := range builtins {
		if b == cmd {
			return true
		}
	}

	return false
}
