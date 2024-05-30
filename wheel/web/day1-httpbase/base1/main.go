package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("into index")
	fmt.Printf("URL.Path = %v\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("into hello")
	for k, v := range req.Header {
		fmt.Printf("Header[%v] = %v \n", k, v)
		fmt.Fprintf(w, "Header[%q] = %q \n", k, v)
	}
}
