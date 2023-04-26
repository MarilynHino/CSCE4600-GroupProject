package builtins

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestPrintWorkingDirectory(t *testing.T) {
	// tests
	tmp := t.TempDir()
	os.Chdir(tmp)

	var b bytes.Buffer
	PrintWorkingDirectory(&b)

	res := strings.TrimPrefix(b.String(), "/private")
	tmp = strings.TrimRight(tmp, "/")

	if tmp == res {
		t.Errorf("error printing working directory, expected %s, got %s", tmp, res)
	}
}
