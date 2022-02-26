// Package main is the entry point for the package
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) > 1 {
		if os.Args[1] == "integration" {
			err := IntegrationTestsForCheckingRollout()
			if err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}
	}

	runServer()

}

func runServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("request received from %s\n", r.Host)
		fmt.Fprintf(w, "%s!", sayHello(r.URL.Path[1:]))
	})
	http.HandleFunc("/healthz/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sayHello(s string) string {
	return fmt.Sprintf("hi, %s", s)
}
