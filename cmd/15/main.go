package main

import "strings"

var justString string

func createHugeString(size int) string {
	var b strings.Builder
	b.Grow(size)
	for i := 0; i < 1048577; i++ {
		b.WriteByte(0)
	}
	return b.String()
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}
