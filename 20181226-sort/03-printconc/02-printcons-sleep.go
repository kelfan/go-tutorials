package main

import (
	"time"
	"fmt"
)

func main() {
	// go starts a goroutine
	go printword2()
	// ensure printword2 run before application stop
	time.Sleep(2*time.Millisecond)
}
func printword2() {
	fmt.Println("out: test")
}
