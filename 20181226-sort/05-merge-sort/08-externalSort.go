package main

import (
	"os"
	"bufio"
	"./pipeline"
	"fmt"
)

func main() {
	p := createPipeline("small.in", 512, 4)
	writeToFile(p, "small.out")
	printFile("small.out")
}
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSourceChunk(file, -1)
	for v := range p {
		fmt.Println(v)
	}
}
func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriterSink(writer, p)
}
func createPipeline(filename string, fileSize, chunkCount int) <-chan int {
	sortResult := []<-chan int{}
	// the checking of that division should be Integer is ignored here
	chunkSize := fileSize / chunkCount
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0) // from 0
		source := pipeline.ReaderSourceChunk(bufio.NewReader(file), chunkSize)

		sortResult = append(sortResult, pipeline.InMemSort(source))

	}
	return pipeline.MergeN(sortResult...)
}
