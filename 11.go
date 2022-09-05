// Реализовать пересечение двух неупорядоченных множеств.

package task

// Intersection проходит по множеству m1 ключи которого дженерик T
// и если данный элемент будет в m2 записывает его в пересечение двух множеств
func Intersection[T comparable](m1, m2 map[T]struct{}) map[T]struct{} {
	m := make(map[T]struct{})
	for k, _ := range m1 {
		if _, ok := m2[k]; ok {
			m[k] = struct{}{}
		}
	}
	return m
}
