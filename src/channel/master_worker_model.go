// 任务和worker的新模型

// 传统的模型是对任务池进行加锁,新模式使用一个通道接收需要处理的任务,一个通道接收处理完成的任务

package main

import (
	"fmt"
	"strconv"
)

const N = 5

func main() {
	pending, done := make(chan * Task), make(chan *Task)
	go sendWork(pending)
	for i := 0; i < N; i++ {
		go Worker(pending, done)
	}
	consumeWork(done)
}

func Worker(in, out chan *Task) {
	for {
		if t, ok := <-in; ok {
			process(t)
			out <- t
		} else {
			close(out)
			break
		}
	}
}

// fake code

type Task struct {
	content string
}

func sendWork(out chan *Task) {
	go func() {
		i := 1
		for k := 0; k < 15; k++{
			str := "Task" + strconv.Itoa(i)
			i++
			var t Task
			t.content = str
			out <- &t
		}
		close(out)
	}()
}

func process(t *Task) *Task {
	t.content += "done"
	return t
}

func consumeWork(in chan *Task) {
	for t := range in {
		fmt.Println(t.content)
	}
}
