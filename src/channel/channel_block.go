package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	go pump(ch1)
	// fmt.Println(<-ch1)
	suck(ch1)
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}