package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func main() {
	html := `<body>

				<div lang="ZH">DIV1</div>
				<div lang="zh-cn">DIV2</div>
				<div lang="en">DIV3</div>
				<span>
					<div>DIV4</div>
				</span>

			</body>
			`

	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))
	if err!=nil{
		log.Fatalln(err)
	}

	// get only direct child div
	dom.Find("body>div").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

	fmt.Print("----------\n")

	// get all under divs
	dom.Find("body div").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}
