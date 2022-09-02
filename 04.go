// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные
// данные из канала и выводят в stdout. Необходима возможность
// выбора количества воркеров при старте.
//
// Программа должна завершаться по нажатию Ctrl+C.
// Выбрать и обосновать способ завершения работы всех воркеров.

package task

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func worker(id int, ctx context.Context, jobs <-chan int) {
	for {
		select {
		case job := <-jobs:
			fmt.Printf("воркер id:%d выполнил работу #%d\n", id, job)
		case <-ctx.Done():
			fmt.Printf("оркер id:%d отменен. детали: %v\n", id, ctx.Err())
			return
		}
	}
}

type WorkerPool struct {
	count int
	jobs  chan int
	Done  chan struct{}
}

func New(count int) WorkerPool {
	return WorkerPool{
		count,
		make(chan int, count),
		make(chan struct{}),
	}
}

func (wp WorkerPool) GenerateJobs(c <-chan os.Signal, cancel context.CancelFunc) {
	for i := 0; ; i++ {
		// Просили, чтобы генератор работал в главом потоке, поэтому поэтому следим за c
		// а пока не следим делаем джоб
		select {
		case <-c:
			cancel()
			return
		default:
			wp.jobs <- i
			time.Sleep(time.Second)
		}
	}
}

func (wp WorkerPool) Run(ctx context.Context) {
	for i := 0; i < wp.count; i++ {
		go worker(i, ctx, wp.jobs)
	}
}

func ReadDataFromMainThread(countWorkers int) {
	wp := New(countWorkers)
	ctx := context.Background()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)

	go wp.Run(ctx)

	wp.GenerateJobs(c, cancel)

	// Второй сигнал означает жёсткий выход из программы
	<-c
	// exit code interrupt
	os.Exit(2)
}
