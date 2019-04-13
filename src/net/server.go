// 简单的tcp服务器

package main

import (
	"fmt"
	"net"
	"strings"
)

var stop bool = false

func main() {
	fmt.Println("Starting the server...")
	// 创建listener
	listener, err := net.Listen("tcp", "localhost:50000")
	if checkError(err, "listening") {
		return
	}

	for !stop {
		conn, err := listener.Accept()
		if checkError(err, "Accept") {
			return 
		}
		go doServerStuff(conn)
	}

	fmt.Printf("Server out")
}

func doServerStuff(conn net.Conn) {
	for !stop {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		// checkError(err, "reading") //即使出错了,也不退出循环
		if checkError(err, "reading") {
			return
		}

		strs := strings.Fields(string(buf[:len]))
		if len > 3 && strs[2] == "SH" {
			stop = true
			break
		}

		fmt.Printf("Received data: %v\n", string(buf[:len]))
	}
	
	conn.Close()
	fmt.Println("goroutine out")
}

func checkError(err error, errStr string) bool {
	if err != nil {
		fmt.Println("Error", errStr, err.Error())
		return true
	}
	return false
}