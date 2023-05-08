package builtins

import (
	"testing"
)

func TestSaveHistory(t *testing.T) {
	history = []string{}

	// Test the SaveHistory function with multiple arguments
	if err := SaveHistory("command1", "command2", "command3"); err != nil {
		t.Errorf("SaveHistory() error = %v, want nil", err)
	}

	// Test that the history has been updated correctly
	if len(history) != 3 {
		t.Errorf("SaveHistory() len(history) = %d, want 3", len(history))
	}

	// Test the SaveHistory function with no arguments
	if err := SaveHistory(); err != nil {
		t.Errorf("SaveHistory() error = %v, want nil", err)
	}

	// Test that the history has not been updated
	if len(history) != 3 {
		t.Errorf("SaveHistory() len(history) = %d, want 3", len(history))
	}
}

func TestPrintHistory(t *testing.T) {
	history = []string{"command1", "command2", "command3"}

	// Test the PrintHistory function
	if err := PrintHistory(); err != nil {
		t.Errorf("PrintHistory() error = %v, want nil", err)
	}
}
