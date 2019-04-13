// 扩展http_fetch.go使之可以从控制台读取url
package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bufio"
	"os"
	"log"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	url, _ := rd.ReadString('\n')
	url = url[:len(url)-1] //去掉最后的换行符
	res, err := http.Get(url)	
	checkError(err)
	data, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Printf("Got: %q", string(data))
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}


