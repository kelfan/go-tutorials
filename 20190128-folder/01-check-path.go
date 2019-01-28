package main

import (
	"fmt"
	"os"
	"time"
)

func CheckFolder(path string) {
	permissionBits := 0700
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.FileMode(permissionBits))
		if err != nil {
			fmt.Println("can't create path: "+path, err)
		}
		fmt.Println("filer.CheckFolder: create file success - " + path)
	} else {
		fmt.Println("filer.CheckFolder: file already exist - " + path)
	}
}

func main() {
	pageTime := time.Now()
	CheckFolder("./Archive/" + pageTime.Format("2006/200601/20060102"))
}
