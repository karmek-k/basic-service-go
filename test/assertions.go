package test

import "testing"

// AssertEqualInts fails if two provided integers are not equal
func AssertEqualInts(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}