// Дана переменная int64.
// Разработать программу которая устанавливает i-й бит в 1 или 0

package task

func SetBit(n int64, pos uint8, one bool) int64 {
	if one {
		return n | (1 << pos)
	} else {
		return n & ^(1 << pos)
	}
}
