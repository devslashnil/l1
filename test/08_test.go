package test

import (
	. "l1"
	"testing"
)

func Test08(t *testing.T) {
	n := SetBit(0b11111111, 3, false)
	var want int64 = 0b11110111
	if n != want {
		t.Errorf("SetBit return expected: %b got: %b", want, n)
	}
}

func Test08_02(t *testing.T) {
	n := SetBit(0b00000001, 4, true)
	var want int64 = 0b00010001
	if n != want {
		t.Errorf("SetBit return expected: %b got: %b", want, n)
	}
}
