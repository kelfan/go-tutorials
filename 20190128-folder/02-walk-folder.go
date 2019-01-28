package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	err := filepath.Walk("data", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			fmt.Println("folder:", info.Name())
		} else {
			fmt.Println("file:", info.Name())
			fmt.Println("path:", path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error in walking path")
	}
	fmt.Println(filepath.Dir("C:/2017/test.html"))
}
