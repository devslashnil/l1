package test

import (
	. "l1"
	"reflect"
	"testing"
)

func Test01(t *testing.T) {
	a := Action{}
	m := "Foo"
	st := reflect.TypeOf(a)
	_, ok := st.MethodByName(m)
	if !ok {
		t.Errorf("%T: don't have %s method", a, m)
	}
}
