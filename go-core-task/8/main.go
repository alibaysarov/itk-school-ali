package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type CustomWaitGroup struct {
	counter int64
	sema    chan struct{}
}

func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		counter: 0,
		sema:    make(chan struct{}, 1),
	}
}

func (wg *CustomWaitGroup) Add(delta int) {
	atomic.AddInt64(&wg.counter, int64(delta))
}

func (wg *CustomWaitGroup) Done() {
	if atomic.AddInt64(&wg.counter, -1) == 0 {

		select {
		case wg.sema <- struct{}{}:
		default:
		}
	}
}

func (wg *CustomWaitGroup) Wait() {
	if atomic.LoadInt64(&wg.counter) == 0 {
		return // ничего ждать не нужно
	}

	<-wg.sema
	wg.sema <- struct{}{}
}

func main() {
	const N = 5
	wg := NewCustomWaitGroup()

	wg.Add(N)

	for i := 0; i < N; i++ {
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(id) * 100 * time.Millisecond)
			fmt.Printf("Горутина %d завершилась\n", id)
		}(i)
	}

	fmt.Println("Ожидание завершения всех горутин...")
	wg.Wait()
	fmt.Println("Все горутины завершены!")
}
