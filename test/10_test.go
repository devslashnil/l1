package test

import (
	. "l1"
	"reflect"
	"testing"
)

func Test10(t *testing.T) {
	r := Bucket([]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5})
	want := map[int][]float64{-20: {-25.0, -27.0, -21.0}, 10: {13.0, 19.0, 15.5}, 20: {24.5}}
	if reflect.DeepEqual(r, want) {
		t.Errorf("SetBit return expected: %v got: %v", want, r)
	}
}
