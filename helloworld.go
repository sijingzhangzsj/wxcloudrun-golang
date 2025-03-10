package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world received a request.")
	target := "Welcome to CloudBase"
	fmt.Fprintf(w, "Hello, %s!\n", target)
}

func main() {
	log.Print("Hello world sample started.")
	http.HandleFunc("/", handler)


	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", 8081), nil))
}
