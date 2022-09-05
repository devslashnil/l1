package test

import (
	. "l1"
	"strings"
	"testing"
)

func Test15(t *testing.T) {
	r := RunSomeFunc()
	want := strings.Repeat("语", 100)
	if r != want {
		t.Errorf("RunSomeFunc return expected: %v got: %v", want, r)
	}
}
