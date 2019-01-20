package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func main() {
	html := `<body>

				<div id="div1">DIV1</div>
				<div>DIV2</div>
				<span>SPAN</span>

			</body>
			`

	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))
	if err!=nil{
		log.Fatalln(err)
	}

	dom.Find("#div1").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}
