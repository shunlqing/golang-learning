package main

import (
	"fmt"
	"net/http"
)

var urls = []string {
	"http://www.baidu.com/",
	"http://golang.org/",
	"http://blog.golang.org",
}

func main() {
	// Excute an HTTP HEAD request for all url's
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error: ", url, err)
		}
		fmt.Println(url, ": ", resp.Status)
	}
}

