package main

import (
	"os"
	"bufio"
	"./pipeline"
	"fmt"
)

func main() {
	p := createPipeline3("large.in", 800000, 4) // fileSize need to check the file size in the folder
	writeToFile3(p, "large.out")
	printFile3("large.out")
}
func printFile3(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSourceBuffer(file, -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}
func writeToFile3(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriterSink(writer, p)
}
func createPipeline3(filename string, fileSize, chunkCount int) <-chan int {
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
		source := pipeline.ReaderSourceBuffer(bufio.NewReader(file), chunkSize)

		sortResult = append(sortResult, pipeline.InMemSortBuffer(source))

	}
	return pipeline.MergeNBuffer(sortResult...)
}
