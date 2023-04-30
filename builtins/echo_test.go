package builtins_test

import (
	"bytes"
	"testing"

	"github.com/jh125486/CSCE4600/Project2/builtins"
)

func TestEcho(t *testing.T) {
	// Create a buffer to use as the output destination
	buf := new(bytes.Buffer)

	// Call the Echo function with some arguments
	args := []string{"hello", "world"}
	err := builtins.Echo(buf, args)
	if err != nil {
		t.Fatalf("Echo failed with error: %v", err)
	}

	// Verify that the output matches what we expect
	expectedOutput := "hello world\n"
	actualOutput := buf.String()
	if actualOutput != expectedOutput {
		t.Errorf("Echo output was incorrect. Expected: %q, actual: %q", expectedOutput, actualOutput)
	}
}
