package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 500; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		if out, ok := <-ch; ok {
			fmt.Println(out)
		} else {
			break
		}
	}

	// use function
	ch2 := make(chan string)
	go printLine(1000, ch2)

	for {
		if msg, ok2 := <-ch2; ok2 {
			fmt.Println(msg)
		} else {
			break
		}
	}
}

func printLine(num int, ch chan string) <-chan string {
	for j := 0; j < num; j++ {
		ch <- fmt.Sprintln(j, ": -------------")
	}
	close(ch)
	return ch
}
