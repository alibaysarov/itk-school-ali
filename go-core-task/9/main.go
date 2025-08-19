package main

import (
	"fmt"
	"time"
)

func startPipeline(in <-chan uint8, out chan<- float64) {
	defer close(out)

	for num := range in {
		f := float64(num)
		cubed := f * f * f
		out <- cubed
	}
}

func main() {
	in := make(chan uint8)
	out := make(chan float64)

	go startPipeline(in, out)
	go func() {
		defer close(in)

		numbers := []uint8{1, 2, 3, 4, 5, 10, 7}
		for _, n := range numbers {
			fmt.Printf("Отправлено в in: %d\n", n)
			in <- n
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for result := range out {
		fmt.Printf("Получено из out (куб): %.2f\n", result)
	}

	fmt.Println("Конвейер завершил работу.")
}
