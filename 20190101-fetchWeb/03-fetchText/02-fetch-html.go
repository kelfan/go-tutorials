package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("http://www.bom.gov.au/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	scanner := bufio.NewScanner(response.Body)
	scanner.Split(bufio.ScanRunes)
	var buf bytes.Buffer
	for scanner.Scan() {
		buf.WriteString(scanner.Text())
	}
	fmt.Println(buf.String())
}