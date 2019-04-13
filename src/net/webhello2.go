
package main

import (
	"net/http"
	"fmt"
	"log"
	"strings"
)

func handleFunc1(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, " + req.URL.Path[len("/hello/"):])
}

func handleFunc2(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, " + strings.ToUpper(req.URL.Path[len("/shouthello/"):]))
}

func main() {
	http.HandleFunc("/hello/", handleFunc1)
	http.HandleFunc("/shouthello/", handleFunc2)
	err := http.ListenAndServe("localhost:9999", nil) 
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}