// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package task

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type CounterAtomic struct {
	count int32
	wg    *sync.WaitGroup
}

func (c *CounterAtomic) Inc() int32 {
	c.wg.Add(1)
	defer c.wg.Done()
	return atomic.AddInt32(&c.count, 1)
}

func (c *CounterAtomic) Get() int32 {
	c.wg.Add(1)
	defer c.wg.Done()
	return atomic.LoadInt32(&c.count)
}

func (c *CounterAtomic) FinalGet() int32 {
	c.wg.Wait()
	return atomic.LoadInt32(&c.count)
}

type CounterMutex struct {
	count int
	mutex *sync.RWMutex
	wg    *sync.WaitGroup
}

func (c *CounterMutex) Inc() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.wg.Add(1)
	defer c.wg.Done()
	c.count++
}

func (c *CounterMutex) Get() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	c.wg.Add(1)
	defer c.wg.Done()
	return c.count
}

func (c *CounterMutex) FinalGet() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	c.wg.Wait()
	return c.count
}

func RunCounter() {
	wg := sync.WaitGroup{}
	c := CounterAtomic{0, &wg}
	for i := 0; i < 1000; i++ {
		go c.Inc()
	}
	fmt.Println("FinalGet = ", c.FinalGet())
}
