package builtins

import (
	"fmt"
	"io"
	"strings"
)

func Echo(w io.Writer, args []string) error {
	// All args after echo command are merged into a single string
	msg := strings.Join(args, " ")

	// Write the message to io.Writer
	_, err := fmt.Fprintln(w, msg)
	if err != nil {
		return err
	}

	return nil
}
