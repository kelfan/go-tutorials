package main

import (
	"./pipeline"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	createNetworkPipeline4("small.in", 512, 4) // fileSize need to check the file size in the folder
	//writeToFile4(p, "large.out")
	//printFile4("large.out")
	time.Sleep(time.Hour)
}
func printFile4(filename string) {
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
func writeToFile4(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriterSink(writer, p)
}

func createPipeline4(filename string, fileSize, chunkCount int) <-chan int {
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

func createNetworkPipeline4(filename string, fileSize, chunkCount int) <-chan int {
	// the checking of that division should be Integer is ignored here
	chunkSize := fileSize / chunkCount
	pipeline.Init()

	sortAddr := []string{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0) // from 0
		source := pipeline.ReaderSourceBuffer(bufio.NewReader(file), chunkSize)

		addr := ":" + strconv.Itoa(7000 + i)
		pipeline.NetworkSink(addr, pipeline.InMemSortBuffer(source))
		sortAddr = append(sortAddr, addr)

	}

	return nil

	sortResult := []<-chan int{}
	for _, addr := range sortAddr {
		sortResult = append(sortResult, pipeline.NetworkSource(addr))
	}

	return pipeline.MergeNBuffer(sortResult...)
}
