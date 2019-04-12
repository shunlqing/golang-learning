// 惰性生成器

//使用int类型通道实现一个int类型的惰性生成器

package main

import (
	"fmt"
)

var resume chan int

func integers() chan int {
	ch := make(chan int)
	count := 0
	go func() {
		for {
			ch <- count
			count++
		}
	}()
	return ch
}

func generateInteger() int {
	return <- resume
}

func main() {
	resume = integers()
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
	
}



