package test

import (
	. "l1"
	"testing"
)

func Test11(t *testing.T) {
	m1 := map[string]struct{}{
		"hi":  {},
		"buy": {},
	}
	m2 := map[string]struct{}{
		"world": {},
	}
	want := map[string]struct{}{
		"hi":    {},
		"buy":   {},
		"world": {},
	}
	r := Merge[string](m1, m2)
	for k, _ := range r {
		_, ok := want[k]
		if !ok {
			t.Errorf("Merge return expected: %v got: %v", want, r)
		}
	}
	if len(r) != len(want) {
		t.Errorf("Merge sizes unequal: %v got: %v", want, r)
	}
}
