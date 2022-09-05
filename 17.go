// Реализовать бинарный поиск встроенными методами языка.

package task

func Find(n int, cmp func(int) int) (i int, found bool) {
	// Определяем cmp(-1) > 0 and cmp(n) <= 0
	// Инвариант: cmp(i-1) > 0, cmp(j) <= 0
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // делим на 2 i+j
		// i ≤ h < j
		if cmp(h) > 0 {
			i = h + 1 // сохраняет cmp(i-1) > 0
		} else {
			j = h // сохраняет cmp(j) <= 0
		}
	}
	// i == j, cmp(i-1) > 0 и cmp(j) <= 0
	return i, i < n && cmp(i) == 0
}
