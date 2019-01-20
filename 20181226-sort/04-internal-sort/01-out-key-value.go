package main

import (
	"sort"
	"fmt"
)

func main() {
	// Create a slice of int
	a := []int{3,9,4,6,2,1,6,4}
	sort.Ints(a)

	// print key and value
	for i, v := range a {
		fmt.Println(i, v)
	}

	// print only key
	for _, v := range a{
		fmt.Println(v)
	}
}
