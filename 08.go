// Дана переменная int64.
// Разработать программу, которая устанавливает i-й бит в 1 или 0

package task

// SetBit прибавляет i-ый бит или обнуляет в зависимости от one.
func SetBit(n int64, pos uint8, one bool) int64 {
	if one {
		// смещает 1 на pos бит, затем прибавляет его к n
		return n | (1 << pos)
	} else {
		// смещает 1 на pos бит, затем меняем биты, чтобы перемножить n на 11..101..1
		return n & ^(1 << pos)
	}
}

// SetBit2 обнуляет i-ый бит и присваивает ему нужное значение
func SetBit2(n int64, pos uint8, one int64) int64 {
	return n & ^(1<<pos) | (one << pos)
}
