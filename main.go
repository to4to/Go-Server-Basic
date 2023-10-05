package main

import (
	"fmt"
	"net/http"
)

func main() {

	fileerver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileerver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello"helloHandler)

	fmt.Println("Server Started")
}
