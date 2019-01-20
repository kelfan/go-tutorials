package main

import (
	"time"
	"fmt"
)

func main() {
	for i := 0; i < 500; i++ {
		go printword4(i)
	}
	time.Sleep(time.Millisecond)
}
func printword4(num int) {
	// dead loop
	for{
		fmt.Printf("test: %d\n", num)
	}
}
