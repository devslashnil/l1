// Разработать программу, которая будет последовательно
// отправлять значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

package task

import (
	"context"
	"fmt"
	"time"
)

// generate бесконечно отправляет в c числа и ждёт ctx.Done
func generate(ctx context.Context, c chan<- int) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			c <- i
			time.Sleep(time.Second)
		}
	}
}

// consume бесконечно ждёт из канала c числа и ждёт ctx.Done
func consume(ctx context.Context, c <-chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case i := <-c:
			fmt.Println(i)
		}
	}
}

// BackAndForthOut запускает конкурентно generate и consume в главном потоке
func BackAndForthOut(sec time.Duration) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, sec*time.Second)
	c := make(chan int)
	// generate запускается конкурентно чтобы выполннение перешло к consume,
	// и не ждало бесконечной работы цикла генерации
	go generate(ctx, c)
	// consume запускается в главном потоке для ожидания окончания работы через контекст
	consume(ctx, c)
}
