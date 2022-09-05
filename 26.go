// Разработать программу, которая проверяет, что все
// символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.

package task

import (
	"unicode"
)

func IsUnique(s string) bool {
	m := make(map[rune]struct{})
	for _, v := range s {
		k := unicode.ToLower(v)
		if _, ok := m[k]; ok {
			return false
		}
		m[k] = struct{}{}
	}
	return true
}
