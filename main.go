package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/.well-known/matrix/server", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json:w")
		fmt.Fprintf(w, "{ \"m.server\": \"matrix.4amlunch.net:443\" }")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}