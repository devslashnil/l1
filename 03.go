package task

func SquareSum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v * v
	}
	c <- sum // send sum to c
}

// ConcurSquareSum разбивает подсчёт суммы квадратов на gorNum горутин
func ConcurSquareSum(s []int, gorNum int) int {
	// длина слайсов на которых будет конкуретно считаться квадраты
	step := len(s)/gorNum + 1
	c := make(chan int)
	i := 0
	for ; i < gorNum; i += step {
		go SquareSum(s[i:i+step], c)
	}
	// квадраты для оставшихся элементов
	go SquareSum(s[i:], c)
	res := 0
	// ждём ответов от горутин
	for {
		select {
		case sum := <-c:
			res += sum
			// подсчитываем кол-во оставшихся горутин
			gorNum--
			if gorNum < 1 {
				return res
			}
		}
	}
}
