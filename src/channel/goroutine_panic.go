//使用goroutine实现计数

package main

import (
	"fmt"
)

func main() {
	ch := tel()
	for {
		fmt.Printf("Counter is %d\n", <-ch)
	}
}

func tel() chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i < 15; i++ {
			ch <- i
			i++
		}
	}()
	return ch
}

// 引发运行时的panic:当tel内部的goroutine执行完后,主函数的goroutine仍旧等待


