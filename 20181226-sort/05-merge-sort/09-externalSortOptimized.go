package main

import (
	"os"
	"bufio"
	"./pipeline"
	"fmt"
)

func main() {
	p := createPipeline2("large.in", 800000, 4) // fileSize need to check the file size in the folder
	writeToFile2(p, "large.out")
	printFile2("large.out")
}
func printFile2(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSourceChunk(file, -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}
func writeToFile2(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriterSink(writer, p)
}
func createPipeline2(filename string, fileSize, chunkCount int) <-chan int {
	// the checking of that division should be Integer is ignored here
	chunkSize := fileSize / chunkCount
	pipeline.Init()

	sortResult := []<-chan int{}
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
