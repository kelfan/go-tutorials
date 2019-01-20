package main

// https://gorbledlog.appspot.com/article?id=58ed897913e0b39b9698fbb2afa55283

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"../util/filer"
)

func main(){
	resp,err := http.Get("http://www.bom.gov.au/")
	if err != nil{
		//handle error
		fmt.Println(err)
		log.Fatal(err)
	}
	if resp.StatusCode == http.StatusOK{
		fmt.Println(resp.StatusCode)
	}
	defer resp.Body.Close()

	buf := make([]byte, 1024)
	//create file
	fname := "bom.html"
	filer.CFolder(fname)
	f,err1 := os.OpenFile(fname,os.O_RDWR|os.O_CREATE|os.O_APPEND,os.ModePerm)
	if err1 != nil{
		panic(err1)
		return
	}
	defer f.Close()

	for {
		n,_ := resp.Body.Read(buf)
		if 0 == n {break}
		f.WriteString(string(buf[:n]))
	}

}