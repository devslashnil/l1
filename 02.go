// Написать программу, которая конкурентно рассчитает
// значение квадратов чисел взятых из массива (2,4,6,8,10)
// и выведет их квадраты в stdout.

package task

import (
	"fmt"
)

// SquareArr осуществляет последовательный вывод элментов s в Stdout
func SquareArr(s []int) {
	for _, v := range s {
		fmt.Printf("%d\n", v*v)
	}
}

// ConcurSquareOut разбивает подсчёт квадратов на gorNum горутин
func ConcurSquareOut(s []int, gorNum int) {
	// длина слайсов на которых будет конкуретно считаться квадраты
	step := len(s)/gorNum + 1
	i := 0
	for ; i < gorNum; i += step {
		go SquareArr(s[i : i+step])
	}
	// квадраты для оставшихся элементов
	go SquareArr(s[i:])
}
