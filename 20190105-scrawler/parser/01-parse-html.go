package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"../util/filer"
)

func main() {
	html := filer.Readfile("bom.html")

	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))
	if err!=nil{
		log.Fatalln(err)
	}

	// get all under divs
	dom.Find("body").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}
