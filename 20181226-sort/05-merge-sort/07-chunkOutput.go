package main

import "fmt"
import (
	"./pipeline"
	"os"
	"bufio"
)

func main() {
	//mergeDemo()
	const filename = "large.in"
	const n = 100000
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close() // equals to finally close the file

	p := pipeline.RandomSource(n)
	writer := bufio.NewWriter(file) // output file through blocks
	pipeline.WriterSink(writer, p)
	writer.Flush() // output data into file

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReaderSourceChunk(bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}
