package task

func produce(in chan<- int) {
	for i := 0; i < 5; i++ {
		in <- i
	}
}

func multiply(in chan<- int, out <-chan int) {

}

func output(out chan<- int) {

}

func Conveyor() {
	in := make(chan int)
	out := make(chan int)

}
