// Реализовать пересечение двух неупорядоченных множеств.

package task

func Merge[T comparable](m1, m2 map[T]struct{}) map[T]struct{} {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}
