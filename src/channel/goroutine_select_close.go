//使用协程进行计数器

//使用额外的通道quit传递协程退出的信息,并用select监听多个通道
package main

import (
	"fmt"
	"os"
)

func main() {
	ch, quit:= tel()
	for {
		select {
		case i := <- ch:
			fmt.Printf("The counter is at %d\n", i)
		case <-quit:
			os.Exit(0)
		}
	}
}

func tel() (chan int, chan bool) {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for i := 1; i < 15; i++ {
			ch <- i
		}
		quit <- true
	}()
	return ch, quit
}