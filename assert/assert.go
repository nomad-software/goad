package assert

import (
	"testing"

	"github.com/nomad-software/goad/constraint"
)

func Eq[T constraint.BuiltinTypes](t *testing.T, actual T, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("failed asserting '%v' (actual) == '%v' (expected)\n", actual, expected)
	}
}
