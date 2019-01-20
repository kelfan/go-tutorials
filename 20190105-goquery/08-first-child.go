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
				<p>P1</p>
				<div lang="zh-cn">DIV2</div>
				<div lang="en">DIV3</div>
				<span>
					<p>P2</p>
					<div>DIV5</div>
				</span>
				<div></div>

			</body>
			`

	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))
	if err!=nil{
		log.Fatalln(err)
	}

	// first type element of each level
	dom.Find("div:first-of-type").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Html())
	})
}
