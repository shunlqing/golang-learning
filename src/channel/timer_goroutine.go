// 定时器和计时器Ticker

package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1e8) //返回一个定时触发的计时器,每1e8纳秒,通道tick收到一个时间(Time类型)
	boom := time.After(5e8) //返回一个定时器,只触发一次,即5e8纳秒之后,通道boom收到一个时间(Time类型)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}