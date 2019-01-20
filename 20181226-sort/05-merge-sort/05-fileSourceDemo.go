package main

import "fmt"
import (
	"./pipeline"
	"os"
)

func main() {
	//mergeDemo()
	const filename = "small.in"
	const n = 64
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close() // equals to finally close the file

	p := pipeline.RandomSource(n)
	pipeline.WriterSink(file, p)

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReaderSource(file)
	for v := range p {
		fmt.Println(v)
	}
}
func mergeDemo() {
	p := pipeline.Merge(
		pipeline.InMemSort(pipeline.ArraySource(5, 2, 7, 7, 4)),
		pipeline.InMemSort(pipeline.ArraySource(7, 3, 6, 8, 2, 4, 6, 3, 9)))
	for v := range p {
		fmt.Println(v)
	}
}
