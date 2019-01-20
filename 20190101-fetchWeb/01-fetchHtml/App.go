package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

/**
	get html info based on url
	cites: https://golangtc.com/t/56f52ef5b09ecc66b900019a
 */

func CheckErr(err error)  {
	if err != nil {
		panic(err)
	}
}

func main() {
	url := "http://www.baidu.com"

	// get request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("request: ", req)

	// get response
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	log.Println("response: ", resp)

	// get response success
	html := ""
	if resp.StatusCode == 200 {
		robots, err := ioutil.ReadAll(resp.Body);
		CheckErr(err)
		html = string(robots)
	}
	log.Println(html)
}
