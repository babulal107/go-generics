package fuzz_testing

func Equal(a []byte, b []byte) bool {
	// if length is not the same - slices are not equal
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
