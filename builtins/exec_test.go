package builtins_test

import (
	"testing"

	"github.com/jh125486/CSCE4600/Project2/builtins"
)

func TestExecCommand(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "no command specified",
			args:    []string{},
			wantErr: true,
		},
		{
			name:    "valid command",
			args:    []string{"echo", "Hello, World!"},
			wantErr: false,
		},
		{
			name:    "invalid command",
			args:    []string{"nonexistent-command"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := builtins.ExecCommand(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}