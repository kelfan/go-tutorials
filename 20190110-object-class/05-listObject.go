package main

import (
	"fmt"
	"time"
)

func main() {
	v := []chan string{}
	a := make(chan string)
	v = append(v, a)
	go func() {
		for {
			v[0] <- "a"
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for {
			if b, ok := <-v[0]; ok {
				fmt.Println(b)
			}
		}

	}()
	time.Sleep(time.Second * 100)
}
