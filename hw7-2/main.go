package main

import (
	"fmt"
	"math/rand"
)

type minmax struct {
	min int
	max int
}

type randomNums struct {
	min int
	max int
	qty int
}

func first(randoms randomNums, ch1 chan int, ch2 chan minmax, done chan error) {
	for i := 0; i < randoms.qty; i++ {
		random := rand.Intn(randoms.max-randoms.min+1) + randoms.min
		fmt.Println("random:", random)
		ch1 <- random
	}
	close(ch1)

	mm := <-ch2
	fmt.Printf("min: %d, max: %d\n", mm.min, mm.max)
	close(done)
}

func second(randoms randomNums, ch1 chan int, ch2 chan minmax) {
	var tempMm minmax
	tempMm.max = randoms.min
	tempMm.min = randoms.max
	for value := range ch1 {
		if value < tempMm.min {
			tempMm.min = value
		}
		if value > tempMm.max {
			tempMm.max = value
		}
	}
	ch2 <- tempMm
	close(ch2)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan minmax)
	done := make(chan error)
	var randoms randomNums
	randoms.min = 1
	randoms.max = 99
	randoms.qty = 8

	go first(randoms, ch1, ch2, done)
	go second(randoms, ch1, ch2)

	<-done
}
