package main

import (
	"fmt"
	"log"
	"net/http"
)

func wenhook(w http.ResponseWriter, r *http.Request) {
	log.Println("Received webhook...")
	body := r.Body
	fmt.Fprintf(w, "The body is: %s", body)
}

func main() {
	http.HandleFunc("/webhook", wenhook)
	http.ListenAndServe(":8080", nil)
}
