package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Running :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
