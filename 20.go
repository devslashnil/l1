// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

package task

import "strings"

// ReverseWords Решение 1
// Разбиваем строку на слайс слов, меняем, соединяем в одну строку
func ReverseWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

// ReverseWords2 Решение 2
func ReverseWords2(s string) string {
	// Переворачиваем каждое слово
	runes := []rune(s)
	start := 0
	for end := 0; end < len(s); end++ {
		// Если встречен пробел переворачиваем слово до него
		if s[end] == ' ' {
			SubReverse(runes, start, end)
			start = end + 1
		}
	}

	// Переворачиваем последнюю строку
	SubReverse(runes, start, len(s)-1)

	// Переворачиваем всё предложение
	SubReverse(runes, 0, len(s)-1)
	return string(runes)
}

func SubReverse(runes []rune, i, j int) string {
	for ; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
