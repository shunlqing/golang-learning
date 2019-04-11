//通道关闭和检测

//使用close关闭通道
//使用_,ok检测通道是否关闭
//使用for-range自动检测通道是否关闭

package main

import "fmt"
import "time"

func main() {
	getData2(sendData())	
	time.Sleep(1e9)
}

func sendData() chan string {
	ch := make(chan string)
	go func() {
		ch <- "Washington"
		ch <- "Tripoli"
		ch <- "London"
		ch <- "Beijing"
		ch <- "Tokio"
		close(ch) 
	}()
	return ch
}

func getData1(ch chan string) {
	go func() {
		for {
			input, open := <-ch //检测通道是否关闭
			if !open {
				break
			}
			fmt.Printf("%s ", input)
		}
	}()
}

func getData2(ch chan string) {
	go func() {
		for input := range ch { //使用for-range自动检测通道是否关闭
			fmt.Printf("%s ", input)
		}
	}()
}