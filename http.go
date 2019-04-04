package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler ...handleFuncを定義するのに必要っぽい
func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	// F は書き込み先を明示的に指定できる
}

//type HandlerFunc func(ResponseWriter, *Request)
func main() {
	http.HandleFunc("/", handler)
	// This tells the http package to handle all requests to the web root("/") with handler.
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// http://localhost:8080/monkeys
