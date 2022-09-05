// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй —
// результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

package task

import "fmt"

func produce(in chan<- int, out chan<- int) {
	for i := 0; i < 5; i++ {
		in <- i
	}
	close(in)
}

func multiply(in <-chan int, out chan<- int) {
	for {
		if n, ok := <-in; ok {
			out <- 2 * n
		} else {
			break
		}
	}
	close(out)
}

func output(out <-chan int) {
	for {
		if n, ok := <-out; ok {
			fmt.Println(n)
		} else {
			break
		}
	}
}

func Conveyor() {
	in := make(chan int)
	out := make(chan int)
	go produce(in, out)
	go multiply(in, out)
	output(out)
}
