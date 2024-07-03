package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello!\n")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("ParseForm() err: ", err)
		return
	}

	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) //this http.FileServer is used to serve static files

	http.Handle("/", fileServer)            //this is used to handle the request
	http.HandleFunc("/hello", helloHandler) //this is also used to handle the request(the difference between Handle and HandleFunc is that HandleFunc is used to handle the request with a function)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Starting server at PORT 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
