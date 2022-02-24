// Package main is the entry point for the package
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s!", sayHello(r.URL.Path[1:]))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sayHello(s string) string {
	return fmt.Sprintf("whats up, %s", s)
}
