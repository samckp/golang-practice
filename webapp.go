package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world from GO Lang!</h1>")
}

func check(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Health Check GO Lang!</h1>")
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/healthcheck", check)
	fmt.Println("Server starting !!")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
