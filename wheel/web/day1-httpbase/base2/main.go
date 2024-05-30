package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Println("into index")
		fmt.Printf("URL.Path = %v\n", r.URL.Path)
	case "/hello":
		fmt.Println("into hello")
		for k, v := range r.Header {
			fmt.Printf("Header[%v] = %v \n", k, v)
			fmt.Fprintf(w, "Header[%q] = %q \n", k, v)
		}
	default:
		fmt.Println("default")
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
