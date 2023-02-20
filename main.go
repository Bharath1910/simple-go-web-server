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

	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit form handler!")
	err := r.ParseForm()
	if err != nil {
		fmt.Println("in error")
		fmt.Println(err)
		return
	}

	fmt.Println("POST Success!")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Println("Name:", name)
	fmt.Println("Address:", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Starting server at port 6060")

	err := http.ListenAndServe(":6060", nil)
	if err != nil {
		log.Fatal(err)
	}
}
