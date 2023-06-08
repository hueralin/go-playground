package simplemath

import "testing"

func TestSqrt(t *testing.T) {
	r := Sqrt(9)
	if r != 3 {
		t.Errorf("Sqrt(9) Failed. Got %d, expected 3", r)
	}
}
