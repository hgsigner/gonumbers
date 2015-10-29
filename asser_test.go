package gonumbers_test

import "testing"

func assert(t *testing.T, expected, got string) {
	if got != expected {
		t.Errorf("\nExpected: %s\nGot:      %s", expected, got)
	}
}
