//创建一个空结构hello并使它实现http.Handler.运行并测试

package main

import (
	"fmt"
	"net/http"
	"log"
)

type hello struct {

}

func (h hello) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, " + req.URL.Path[1:])
}

func main() {
	http.Handle("/", hello{})
	err := http.ListenAndServe("localhost:9090", nil)
	if err != nil {
		log.Println("ListenAndServe: ", err.Error())
	}
}