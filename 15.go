//К каким негативным последствиям может привести
//данный фрагмент кода, и как это исправить?
//Приведите корректный пример реализации.
//
//var justString string
//func someFunc() {
//v := createHugeString(1 << 10)
//justString = v[:100]
//}
//
//func main() {
//someFunc()
//}

package task

import "strings"

var justString string

// createHugeString создает строку длины count из символов 语
func createHugeString(count int) string {
	return strings.Repeat("语", count)
}

// someFunc прообразует байтовый слайс в рунный
func someFunc() {
	v := createHugeString(1 << 10)
	// In Go, a string is in effect a read-only slice of bytes.
	// байта не достаточно для юникод символов
	s := []rune(v)
	justString = string(s[:100])
}

// RunSomeFunc запускает программу
func RunSomeFunc() string {
	someFunc()
	return justString
}
