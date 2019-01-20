package main

import (
	"fmt"
	"io/ioutil"
)

// https://golangbot.com/read-files/

func main() {
	data, err := ioutil.ReadFile("bom.html")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
