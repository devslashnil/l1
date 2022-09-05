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

// worker бесконечно ожидает работу из jobs либо отмену из ctx.Done
func worker(id int, ctx context.Context, jobs <-chan int) {
	for {
		select {
		// получение работы
		case job := <-jobs:
			fmt.Printf("воркер id:%d выполнил работу #%d\n", id, job)
		// ожидание отмены контекстом
		case <-ctx.Done():
			fmt.Printf("оркер id:%d отменен. детали: %v\n", id, ctx.Err())
			return
		}
	}
}

type WorkerPool struct {
	count int // кол-во воркеров в пуле
	// буфферезированный канал в который поступает работа для воркеров
	// размер буффера равен кол-во воркеров
	// это позволяет не блокировать отправку в канал до count элементов
	// и не блокировать полчуение из канала пока он не пустой
	// благодаря этому оптимизируется работа по готовности джобов и воркеров
	jobs chan int
}

// New конструктор для WorkerPool
func New(count int) WorkerPool {
	return WorkerPool{
		count,
		// буферизированный канал размера count
		make(chan int, count),
	}
}

// GenerateJobs генерит бесконечное число работ и ждёт интерапта, чтобы заканцелить всех воркеров
func (wp WorkerPool) GenerateJobs(c <-chan os.Signal, cancel context.CancelFunc) {
	for i := 0; ; i++ {
		select {
		// ждём сигнала об интерапте
		case <-c:
			// отправляем в канал ctx.Done для отмены работы воркеров
			cancel()
			return
		// пока сигнала об интерапте нет, отправляем работы в канал раз в секунду для наглядности
		default:
			wp.jobs <- i
			// сон для наглядности
			time.Sleep(time.Second)
		}
	}
}

// Run запускает конкуретную работу из wp.count воркеров
// передаёт им конекст который задает условие остановки работы воркеров
func (wp WorkerPool) Run(ctx context.Context) {
	for i := 0; i < wp.count; i++ {
		// конкурентый старт
		go worker(i, ctx, wp.jobs)
	}
}

// ReadDataFromMainThread инициализирует переменные и запускает воркеров
// также ожидает второго интерапта, чтобы остановить программу
func ReadDataFromMainThread(countWorkers int) {
	wp := New(countWorkers)
	ctx := context.Background()
	c := make(chan os.Signal, 1)
	// направляем os.Interrupt в с
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)

	// конкурентно запускаем воркеров
	go wp.Run(ctx)

	// Генератор работает в главом потоке, т.к. этого требует условие задачи
	wp.GenerateJobs(c, cancel)

	// Ждём второго сигнала os.Interrupt
	// Второй сигнал означает жёсткий выход из программы
	<-c
	// Выходим с exit code interrupt
	os.Exit(2)
}
