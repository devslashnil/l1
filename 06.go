package task

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 1) Горутина ждёт от безысходности
func goroutine1(c <-chan int) {
	fmt.Println("goroutine1 alive!")

	// просто уснула
	time.Sleep(1)

	// получение из канала откуда ничего не пришлют
	<-make(chan int)

	// отправка в nil канал
	(chan int)(nil) <- 0

	// лок на уже залоконом мютексе
	mu := sync.Mutex{}
	mu.Lock()
	mu.Lock()

	// пустой селект
	select {}
}

// 2) Горутина ждёт, чтобы получить из канала
func goroutine2(c <-chan int) {
	fmt.Println("goroutine2 alive!")
	<-c
	fmt.Println("goroutine2 has passed away...")
}

// 3) Горутина ждёт,чтобы отправить в канал
func goroutine3(c chan<- int) {
	fmt.Println("goroutine3 alive!")
	c <- 0
	fmt.Println("goroutine3 has passed away...")
}

// 4) Горутина ждёт, чтобы получить из канала из контекста
func goroutine4(ctx context.Context) {
	fmt.Println("goroutine4 alive!")
	<-ctx.Done()
	fmt.Println("goroutine4 has passed away...")
}

// 5) Горутина ждёт группу ожидания
func goroutine5(wg *sync.WaitGroup) {
	fmt.Println("goroutine5 alive!")
	wg.Wait()
	fmt.Println("goroutine5 has passed away...")
}

// 6) Горутина ждёт когда её разлочат
func goroutine6(mu *sync.Mutex) {
	fmt.Println("goroutine6 alive!")
	mu.Lock()
	fmt.Println("goroutine6 has passed away...")
}

// 7) Горутина ждёт Cond
func goroutine7(b *bool, c *sync.Cond) {
	fmt.Println("goroutine7 alive!")
	c.L.Lock()
	for *b {
		c.Wait()
	}
	c.L.Unlock()
	fmt.Println("goroutine7 has passed away...")
}

// 8) Горутина ждёт переменную
func goroutine8(b *bool) {
	fmt.Println("goroutine8 alive!")
	for {
		if *b {
			break
		}
	}
	fmt.Println("goroutine8 has passed away...")
}

// Way1 пример ждущей горутины
func Way1() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	go goroutine4(ctx)
}
