package test

import (
	. "l1"
	"testing"
)

func Test03(t *testing.T) {
	s := []int{2, 4, 6, 8, 10}
	sum := ConcurSquareSum(s, 2)
	if sum != 220 {
		t.Errorf("ConcurSquareSum return expected 220 got %d", sum)
	}
}
