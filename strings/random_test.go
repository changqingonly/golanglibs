package golang_strings

import (
	"testing"
)

// go test -v -run=^TestRandStr$
func TestRandStr(t *testing.T) {
	if len(RandomStr(30)) != 30 {
		t.Error("e")
	}
}
