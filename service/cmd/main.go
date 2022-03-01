// Package main is the entry point for the package
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("stefan", os.Args)

	if len(os.Args) > 1 {
		fmt.Println(os.Args[2])
		if os.Args[1] == "integration" {
			address := os.Args[2]
			err := IntegrationTestsForCheckingRollout(address)
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
		fmt.Fprintf(w, "hello %s!", sayHello(r.URL.Path[1:]))
	})
	http.HandleFunc("/healthz/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("starting server")
	log.Fatal(http.ListenAndServe(":7836", nil))
}

func sayHello(s string) string {
	return fmt.Sprintf("hello, %s", s)
}
