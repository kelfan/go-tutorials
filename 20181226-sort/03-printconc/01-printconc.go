package main

import "fmt"

func main() {
	// no output because the application stop before running the printword function
	go printword()
}
func printword() {
	fmt.Println("out: test")
}
