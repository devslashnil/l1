package test

import (
	. "l1"
	"testing"
)

func Test08(t *testing.T) {
	n := SetBit(0b11111111, 4, false)
	if n != 0b11110111 {
		t.Errorf("SetBit return expected: %b got: %b", 0b11110111, n)
	}
}
