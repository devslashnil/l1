package task

import (
	"context"
	"fmt"
	"time"
)

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

func BackAndForthOut(sec time.Duration) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, sec*time.Second)
	c := make(chan int)
	go generate(ctx, c)
	consume(ctx, c)

}
