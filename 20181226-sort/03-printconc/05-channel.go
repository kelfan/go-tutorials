package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	for i := 0; i < 500; i++ {
		go printword5(i, ch)
	}

	// dead loop and need stop manually
	// output result
	for{
		msg := <- ch
		fmt.Println(msg)
	}

	//time.Sleep(time.Millisecond)
}
func printword5(num int, ch chan string) {
	// dead loop
	for{
		ch <- fmt.Sprintf("test: %d\n", num)
	}
}
