// 协程恢复

package main

import (
	"fmt"
	"strconv"
	"time"
	"math/rand"
)

func server(workChan <-chan *Work) {
	for work := range workChan {
		go safelyDo(work)
	}
}

//恢复协程的惯用模式
func safelyDo(work *Work) {
	defer func() {
		if err := recover(); err != nil {
			//log.Printf("Work failed with %s in %v", err, work)
			fmt.Printf("Work failed with %s in %v\n", err, work)
		}
	}()

	do(work)
}

// fake code
type Work struct {
	content string
}

func do(work *Work) {
	//模拟随机panic
	if rand.Uint32() % 2 == 1 {
		panic("rand error")
	}
	fmt.Printf("do work: %s\n", work.content)
}

func workPoster() chan *Work {
	ch := make(chan *Work)
	go func() {
		for i := 1; i < 20; i++ {
			str := "work_" + strconv.Itoa(i)
			var w Work
			w.content = str
			ch <- &w
		}
		close(ch)
	}()

	return ch
}

func main() {
	server(workPoster())
	time.Sleep(3e9)
}

// recover只有在defer修饰的函数中调用才会返回非空