//打印素数
package main

import (
	"fmt"
)

//发送序列2,3,4,5,6...到通道ch
func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

//拷贝通道in的值到通道out,移除可以被prime整除的值
func filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i % prime != 0 {
			out <- i
		}
	}
}

func main() {
	ch := make(chan int)
	go generate(ch)
	for {
		prime := <-ch
		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}