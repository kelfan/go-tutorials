package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

/**
	cite: https://jiorry.iteye.com/blog/1879144
 */

func main() {
	resp, err := http.Get("http://baidu.com")
	defer resp.Body.Close()

	if err != nil {
		log.Println("error: ", err)
	}else{
		b, _ := ioutil.ReadAll(resp.Body)
		log.Println("body: ", string(b))
	}
}
