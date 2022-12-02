package fuzz_testing

import "testing"

func TestEqual(t *testing.T) {
	if !Equal([]byte{'f', 'u', 'z', 'z'}, []byte{'f', 'u', 'z', 'z'}) {
		t.Error("expected true, got false")
	}
}

func FuzzEqual(f *testing.F) {
	// target, can be only one per test
	// values of a and b will be auto-generated
	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		Equal(a, b)
	})
}
