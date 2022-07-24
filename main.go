package main

import (
	"fmt"
	"log"
	"net/http"
)

func wenhook(w http.ResponseWriter, r *http.Request) {
	log.Println("Received webhook...")
	body := w.Header()
	log.Println("the webhook...: %s", body)
	fmt.Fprintf(w, "The body is: %s", body)
}

// This was a test for github webhook
// I get the printout but not the body
// I don't know how to do it yet....
func main() {
	http.HandleFunc("/webhook", wenhook)
	http.ListenAndServe(":8080", nil)
}
