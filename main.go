package main

import (
	"fmt"
	"net/http"
)

func wenhook(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	fmt.Fprintf(w, "The body is: %s", body)
}

func main() {
	http.HandleFunc("/", wenhook)
	http.ListenAndServe(":8080", nil)
}
