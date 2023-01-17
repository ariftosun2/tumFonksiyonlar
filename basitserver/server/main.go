package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func urlHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("basit servis")
	req, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Opps", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "hello world:%s", req)
}

func main() {
	http.HandleFunc("/", urlHandler)
	(http.ListenAndServe(":8080", nil))
}
