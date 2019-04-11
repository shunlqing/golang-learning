//使用协程进行计数
//使用close和for-range搭配用于通道的关闭和检测
package main

import (
	"fmt"
)

func main() {
	ch := tel()
	for v := range ch {
		fmt.Printf("Counter is %d\n", v)
	}
}

func tel() chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i < 15; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}