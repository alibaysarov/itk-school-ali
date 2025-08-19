package main

import (
	"fmt"
	"sync"
)

func mergeChannels(channels ...<-chan int) <-chan int {
	merged := make(chan int)
	go func() {
		var wg sync.WaitGroup
		for _, ch := range channels {
			wg.Add(1)
			go func(c <-chan int) {
				defer wg.Done()
				for val := range c {
					merged <- val
				}
			}(ch)
		}
		go func() {
			wg.Wait()
			close(merged)
		}()
	}()

	return merged
}

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)
	ch3 := make(chan int, 2)

	// Отправляем данные
	ch1 <- 1
	ch1 <- 2

	ch2 <- 3
	ch2 <- 4

	ch3 <- 5
	ch3 <- 6

	// Закрываем каналы
	close(ch1)
	close(ch2)
	close(ch3)

	// Сливаем каналы
	merged := mergeChannels(ch1, ch2, ch3)

	// Читаем из объединенного канала
	for val := range merged {
		fmt.Println(val)
	}
}
