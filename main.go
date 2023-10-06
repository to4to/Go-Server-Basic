package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method Not Valid", http.StatusNotFound)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse Form() err: %v", err)

	}
	fmt.Fprintf(w, "Post Request Successful ")
	name := r.FormValue("name ")
	address := r.FormValue("address ")

	fmt.Fprintf(w, "name =%s", name)
	fmt.Fprintf(
		w,
		"address is %s",
		address,
	)
}
func main() {

	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server Started at port 8080")

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}

}
