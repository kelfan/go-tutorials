package main

import "fmt"
import "./pipeline"

func main() {
	p := pipeline.ArraySource(5,2,7,7,4)
	for {
		if num, ok := <- p; ok {
			fmt.Println(num)
		}else{
			break
		}
	}
}
