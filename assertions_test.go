package gonumbers_test

import (
	"strings"
	"testing"
)

func assert(t *testing.T, expected, got interface{}) {
	if got != expected {
		t.Fatalf("\nExpected: %v\nGot:      %v", expected, got)
	}
}

func assert_not(t *testing.T, expected, got interface{}) {
	if got == expected {
		t.Fatalf("\nExpected: %v\nGot:      %v", expected, got)
	}
}

func contains(t *testing.T, s, cont string) {
	if !strings.Contains(s, cont) {
		t.Fatalf("%s\nDoes not contain: %s", s, cont)
	}
}

func notContains(t *testing.T, s, cont string) {
	if strings.Contains(s, cont) {
		t.Fatalf("%s\nShould not contain: %s", s, cont)
	}
}
