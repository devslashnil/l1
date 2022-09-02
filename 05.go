package task

import (
	"context"
	"time"
)

func generate(ctx context.Context, c chan<- int) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			c <- i
		}
	}
}

func consume(ctx context.Context, c <-chan int) {

}

func BackAndForthOut(sec time.Duration) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, sec*time.Second)

}
