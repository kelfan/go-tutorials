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
					<div>DIV4</div>
				</span>
				<p>P2</p>

			</body>
			`

	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))
	if err!=nil{
		log.Fatalln(err)
	}

	// get the next p
	dom.Find("div[lang=zh]+p").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

	fmt.Println("------------")
	// any element with the same parent
	dom.Find("div[lang=zh]~p").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
	fmt.Println("------------")
	// get element contains DIV2
	dom.Find("div:contains(DIV2)").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
	fmt.Println("------------")

	// 筛选出包含div元素的span节点
	dom.Find("span:has(div)").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}
