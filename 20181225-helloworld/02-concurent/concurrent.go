package main

import (
	"fmt"
	"time"
)

func out(line int) {
	i := 0
	for {
		fmt.Println("line ",line, ": ", i)
		i++
	}
}

func main() {
	for i := 0; i<=10; i++ {
		go out(i)
	}
	time.Sleep(time.Microsecond * 1000)
}
