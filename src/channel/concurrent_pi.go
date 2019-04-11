// 使用公式计算pi的近似值

package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// ch := formula()
	// neg := true
	// var pi float64 = 0.0
	// for {
	// 	v := <-ch
	// 	if neg {
	// 		pi += v
	// 	} else {
	// 		pi -= v
	// 	}
	// 	neg = !neg
	// 	fmt.Println(pi)
	// }

	start := time.Now()
	fmt.Println(CalculatePi(5000))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

//不靠谱版本,非常慢

func formula() chan float64 {
	ch := make(chan float64)
	go func() {
		i := 1.0
		for {
			ch <- 4 / i
			i += 2
		}
	}()
	return ch
}

// 靠谱版本
// 同时开启n个协程计算每一项的值
func CalculatePi(n int) float64 {
	ch := make(chan float64)
	for i := 0; i < n; i++ {
		go term(ch, float64(i))
	}
	f := 0.0

	for i := 0; i < n; i++ {
		f += <-ch
	}
	return f
}

func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2 * k + 1)
}