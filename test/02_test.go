package test

import (
	. "l1"
	"testing"
)

func Test02(t *testing.T) {
	s := []int{2, 4, 6, 8, 10}
	ConcurSquareOut(s, 2)
}
