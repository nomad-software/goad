package assert

import (
	"testing"
)

func Eq[T comparable](t *testing.T, actual T, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("failed asserting %v (actual) == %v (expected)\n", actual, expected)
	}
}
