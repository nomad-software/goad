package assert

import (
	"testing"
)

// Eq is a helper function to test the equality of two similarly typed values.
func Eq[T comparable](t *testing.T, actual T, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("failed asserting %v (actual) == %v (expected)\n", actual, expected)
	}
}
