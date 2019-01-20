package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func main() {
	html := `<body>
				<div lang="zh">DIV1</div>
				<span>
					<div>DIV5</div>
				</span>

			</body>
			`

	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))
	if err!=nil{
		log.Fatalln(err)
	}
	// div or span
	dom.Find("div,span").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Html())
	})
}
