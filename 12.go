// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.

package task

// Subset устанавливает ключи в m равные значениям строк, тем самым убираю дубликаты и создавая множество
func Subset(s ...string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, v := range s {
		m[v] = struct{}{}
	}
	return m
}
