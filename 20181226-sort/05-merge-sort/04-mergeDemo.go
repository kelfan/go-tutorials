package main

import (
	"./pipeline"
	"fmt"
)

func main() {
	p := pipeline.Merge(
		pipeline.InMemSort(pipeline.ArraySource(5, 2, 7, 7, 4)),
		pipeline.InMemSort(pipeline.ArraySource(7, 3, 6, 8, 2, 4, 6, 3, 9)))
	for v := range p {
		fmt.Println(v)
	}
}
