// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй —
// результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

package task

import "fmt"

// produce генерит в in 5 чисел, затем закрывает канал in,
// чтобы другие функции могли узнать о завершении работы канала
func produce(in chan<- int) {
	for i := 0; i < 5; i++ {
		in <- i
	}
	close(in)
}

// multiply ждёт канал in для умножения результата на 2 и отправки в out
// и в случае закрытия in закрывает канал out
// чтобы другие функции могли узнать о завершении работы канала
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

// output ждёт канал out для вывода значений, которые пришли в него
func output(out <-chan int) {
	for {
		if n, ok := <-out; ok {
			fmt.Println(n)
		} else {
			break
		}
	}
}

// Conveyor запускает конвейер функций
func Conveyor() {
	in := make(chan int)
	out := make(chan int)
	// конкурентно запускаем циклы, чтобы очередь выполнения дошла до каждого и не ждала окончание любого из них
	go produce(in)
	go multiply(in, out)
	// output в главном потоке, т.к. ждём завершения других горутин
	output(out)
}
