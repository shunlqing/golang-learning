//做一个随机位生成器,程序提供无限的随机0或1的序列

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch := bitgen2()
	for {
		fmt.Printf("%d ", <-ch)
	}
}

//使用math/rand包
func bitgen1() chan int {
	ch := make(chan int)
	go func() {
		for {
			if rand.Uint32()%2 == 0 {
				ch <- 0
			} else {
				ch <- 1
			}
		}
	}()
	return ch
}

//使用select随机选择就绪的通道的原理
func bitgen2() chan int {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()
	return ch
}