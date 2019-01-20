package pipeline

import (
	"fmt"
	"sort"
	"io"
	"encoding/binary"
	"math/rand"
	"time"
)

var startTime time.Time

func Init(){
	startTime = time.Now()
}

/**
	input ints into channel
 */
func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out) // close channel after input all items
	}()
	return out
}

/**
	sort elements from chan int
 */
func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// read into memory
		a := []int{} // new an empty int Array
		for v := range in { // get each element from input
			a = append(a, v)
		}
		fmt.Println("Read done:", time.Now().Sub(startTime))

		// sort
		sort.Ints(a)
		fmt.Println("InMemSort done:", time.Now().Sub(startTime))


		// output
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

/**
	sort elements from chan int
 */
func InMemSortBuffer(in <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		// read into memory
		a := []int{} // new an empty int Array
		for v := range in { // get each element from input
			a = append(a, v)
		}
		fmt.Println("Read done:", time.Now().Sub(startTime))

		// sort
		sort.Ints(a)
		fmt.Println("InMemSort done:", time.Now().Sub(startTime))


		// output
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

/**
	merge 2 chan ints
	compare the first item of each chan int, and pick the smaller one out
 */
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// pick first item out
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			// !ok2 means there is no items left in in2?
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				// pick the next item out
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
		fmt.Println("Merge done:", time.Now().Sub(startTime))
	}()
	return out
}

/**
	merge 2 chan ints
	compare the first item of each chan int, and pick the smaller one out
 */
func MergeBuffer(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		// pick first item out
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			// !ok2 means there is no items left in in2?
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				// pick the next item out
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
		fmt.Println("Merge done:", time.Now().Sub(startTime))
	}()
	return out
}

/**
	get chan ints from file source
 */
func ReaderSource(reader io.Reader) <-chan int {
	out := make(chan int)
	go func() {
		buffer := make([]byte, 8)
		for {
			// read file block into buffer
			n, err := reader.Read(buffer)
			if n > 0 {
				// transfer 64 bit block into int
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
				// if error is not null
				if err != nil {
					break
				}
			}
		}
		close(out)
	}()
	return out
}

/**
	get chan ints from file source
 */
func ReaderSourceChunk(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int)
	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			// read file block into buffer
			n, err := reader.Read(buffer)
			bytesRead += n
			if n > 0 {
				// transfer 64 bit block into int
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
				// if error is not null or get bytesRead bigger than chunkSize
				if err != nil || (chunkSize != -1 && bytesRead >= chunkSize) {
					break
				}
			}
		}
		close(out)
	}()
	return out
}

/**
	get chan ints from file source with  buffer
 */
func ReaderSourceBuffer(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			// read file block into buffer
			n, err := reader.Read(buffer)
			bytesRead += n
			if n > 0 {
				// transfer 64 bit block into int
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
				// if error is not null or get bytesRead bigger than chunkSize
				if err != nil || (chunkSize != -1 && bytesRead >= chunkSize) {
					break
				}
			}
		}
		close(out)
	}()
	return out
}

/**
	write ints into file
 */
func WriterSink(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

/**
 	get random ints
 */
func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

/**
	merge multiple ints
 */
func MergeN(inputs ... <-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}
	m := len(inputs) / 2
	// merge inputs[0..m) and inputs [m..end)
	return Merge(
		MergeN(inputs[:m]...),
		MergeN(inputs[m:]...))

}

/**
	merge multiple ints
 */
func MergeNBuffer(inputs ... <-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}
	m := len(inputs) / 2
	// merge inputs[0..m) and inputs [m..end)
	return MergeBuffer(
		MergeNBuffer(inputs[:m]...),
		MergeNBuffer(inputs[m:]...))

}