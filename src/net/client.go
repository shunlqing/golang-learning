//

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//打开链接
	conn, err := net.Dial("tcp", "localhost:50000")
	if checkError(err, "dialing") {
		return 
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\n")
	
	for {
		fmt.Println("What to send to the server? Type Q to quit. Type SH to stop server")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")
		if trimmedInput == "Q" {
			conn.Close()
			return 
		}

		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
		if checkError(err, "writing") {
			return
		}
	}
}

func checkError(err error, errStr string) bool {
	if err != nil {
		fmt.Println("Error", errStr, err.Error())
		return true
	}
	return false
}