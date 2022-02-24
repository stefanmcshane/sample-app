// Package main is the entry point for the package
package main

import "fmt"

func main() {
	fmt.Println(sayHello("stefan"))
}

func sayHello(s string) string {
	return fmt.Sprintf("hello, %s", s)
}
