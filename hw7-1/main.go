package main

import (
	"fmt"
	"math/rand"
)

func first(ch1 chan int) {
	for i := 0; i < 8; i++ {
		random := rand.Intn(100)
		fmt.Println("random:", random)
		ch1 <- random
	}
	close(ch1)
}

func second(ch1 chan int, ch2 chan float32) {
	var sum int
	var amount int
	for value := range ch1 {
		sum += value
		amount++
	}
	fmt.Printf("sum/amount: %d/%d\n", sum, amount)
	avg := (float32(sum)) / (float32(amount))
	ch2 <- avg
	close(ch2)
}

func third(ch2 chan float32, done chan error) {
	result := <-ch2
	fmt.Println("avg:", result)
	close(done)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan float32)
	done := make(chan error)
	go first(ch1)
	go second(ch1, ch2)
	go third(ch2, done)

	<-done
}
