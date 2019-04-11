// 使用select监听多个通道

package main

import (
	"fmt"
	"time"
)

func main() {
	suck(pump1(), pump2())
	time.Sleep(1e9)
}

func pump1() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i * 2
			time.Sleep(1e9)
		}
	}()

	return ch
}

func pump2() chan int {
	ch := make(chan int) 
	go func() {
		for i := 0; ; i++ {
			ch <- i + 5
			time.Sleep(2e9)
		}
	}()
	return ch
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <- ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <- ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		// default:
		// 	fmt.Printf("default...\n")
		}
	}
}