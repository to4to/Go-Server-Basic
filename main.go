package main

import "net/http"

func main() {

	fileerver := http.FileServer(http.Dir("./static"))
	http.Handle()
}
