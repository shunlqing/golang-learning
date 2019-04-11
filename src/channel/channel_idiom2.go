//通道惯用方式: 使用for-range循环获取通道中的值

package main

import (
	"fmt"
	"time"
)

func main() {
	suck(pump())
	time.Sleep(1e9)
}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	go func() {
		for v := range ch { //使用for-range循环从通道中获取值
			fmt.Println(v)
		}
	}()
}