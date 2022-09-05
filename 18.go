// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package task

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type CounterAtomic int32

func (c *CounterAtomic) Inc() int32 {
	return atomic.AddInt32((*int32)(c), 1)
}

func (c *CounterAtomic) Get() int32 {
	return atomic.LoadInt32((*int32)(c))
}

type CounterMutex struct {
	count int
	mutex *sync.RWMutex
}

func (c *CounterMutex) Inc() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count++
}

func (c *CounterMutex) Get() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.count
}

// RunCounter запускает подсчет n инкриметов
func RunCounter() {
	wg := sync.WaitGroup{}
	c := CounterAtomic(0)
	for i := 0; i < 10000000; i++ {
		wg.Add(1)
		go func() {
			// Уменьшаем счетчик после завершения горутины.
			defer wg.Done()
			c.Inc()
		}()
	}
	// ждём завершения всех горутин
	wg.Wait()
	fmt.Println("FinalGet = ", c)
}
