package main

import (
	"./pipeline"
	"fmt"
)

func main() {
	p := pipeline.ArraySource(5,2,7,7,4)
	for v := range p{
		fmt.Println(v)
	}
}
