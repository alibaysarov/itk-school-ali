package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(ch chan int) {
	for i := 0; i <= 5; i++ {

		fmt.Printf("Отправка данных %d\n", i)
		ch <- rand.Int()
		fmt.Printf("Данные %d ушли\n", i)
	}
}
func consumer(ch chan int) {

	for a := range ch {
		fmt.Printf("Данные пришли %d\n", a)
		time.Sleep(200 * time.Millisecond)
	}

}
func main() {

	ch := make(chan int) // небуферизированный

	// Ждём завершения
	go producer(ch)
	go consumer(ch)
	time.Sleep(5 * time.Second)

	fmt.Println("Готово.")

}
