//使用公式计算Pi的近似值

// 使用GOMAXPROCS,开启和GOMAXPROCS同样多的协程

package main

import (
	"math"
	"time"
	"fmt"
	"runtime"
)

const NCPU = 2

func main() {
	start := time.Now()
	runtime.GOMAXPROCS(2) //设置CPU的核心数
	fmt.Println(CalculatePi(5000))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func CalculatePi(end int) float64 {
	ch := make(chan float64)
	for i := 0; i < NCPU; i++ {
		go term(ch, i * end/NCPU, (i+1) * end / NCPU)
	}
	result := 0.0
	for i := 0; i < NCPU; i++ {
		result += <-ch
	}

	return result
}

func term(ch chan float64, start, end int) {
	result := 0.0
	for i := start; i < end; i++ {
		x := float64(i)
		result += 4 * (math.Pow(-1, x) / (2.0 * x + 1.0))
	}
	ch <- result
}