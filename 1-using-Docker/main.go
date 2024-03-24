package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	log.Println("Starting: v0.3 of Building Apps For K8s app")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	log.Printf("Request received from %s", r.RemoteAddr)
	fmt.Fprintf(w, "v0.3 Building Apps For K8s app says Hi")
}
