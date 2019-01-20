package main

import (
	"time"
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		go printword3(i)
	}
	time.Sleep(time.Millisecond)
}
func printword3(num int) {
	fmt.Printf("test: %d\n", num)
}
