// Удалить i-ый элемент из слайса.

package task

// Remove удаляет с сохранением порядка, соединяя слайсы между элементов
func Remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

// RemoveOrderless удаляет без сохранения порядка
// меняем i элемент на последний и возвращает слайс длины len(s)-1
// выигрыш: не нужно двигать элементы в памяти
func RemoveOrderless[T any](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
