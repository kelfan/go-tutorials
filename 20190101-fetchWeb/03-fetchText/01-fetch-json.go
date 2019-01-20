package main

// https://stackoverflow.com/questions/38807903/how-do-i-handle-plain-text-http-get-response-in-golang

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "http://country.io/capital.json"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)

	fmt.Println(responseString)
}