package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "version 1.0\n")
}

func main() {
	fmt.Println("Serving on port 8000")
	http.HandleFunc("/version", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
