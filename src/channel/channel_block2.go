//写一个通道证明它的阻塞性,开启一个协程接收通道的数据,持续15秒,然后给通道写入一个值

package main

import (
	"fmt"
	"time"
)

func main() {
	// var ch1 chan int 不能仅仅声明
	ch1 := make(chan int)  

	fmt.Println("go suck...")
	go func() {
		fmt.Println("suck ...")
		time.Sleep(15 * 1e9)
		fmt.Println(<-ch1)
	}()
	ch1 <- 11
	fmt.Println("send finish.")
}

