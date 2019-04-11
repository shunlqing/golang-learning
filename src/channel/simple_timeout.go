// 简单超时模式

//创建一个定时器并获取一个通道,使用select同时监听

package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := createTimerout(1e9)
	ch := make(chan int)
	select {
	case <-ch:
		// ...
	case <-timeout:
		fmt.Println("the read from ch has timed out")
		break
	case <-time.After(1e8): //使用time.After完成超时功能
		fmt.Println("timed out (use time.After)")
		break
	}
}

//创建一个定时器,夷?这不是time.After就能完成的吗?
func createTimerout(d int64) chan bool {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Duration(d))
		timeout <- true
	}()
	return timeout
}

/*
 * 注意点:
 * 1. 设置缓冲大小为1是必要的,可以避免协程死锁以及确保超时的通道可以被垃圾回收
 * 2. select对case的选择是伪随机的
*/