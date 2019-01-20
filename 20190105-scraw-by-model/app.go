package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
)

//https://www.jianshu.com/p/99456156e1ce
type SchoolObj struct {
	rankTypeName string
	RankIndex int
	SchoolName string
	EnrollOrder string
	StarLevel string
	LocationName string
	SchoolType  string
	UrlAddress string
	SchoolTags []string
}




func GaokaoquanRank(urlAddress string) []SchoolObj {

	// Request the HTML page.
	res, err := http.Get(urlAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}


	var array [] SchoolObj
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".bangTable table tr").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		var obj SchoolObj
		obj.RankIndex, _ = strconv.Atoi(s.Find(".t1 span").Text())
		obj.SchoolName = s.Find(".t2 a").Text()
		obj.UrlAddress ,_ = s.Find(".t2 a").Attr("href")
		obj.LocationName = s.Find(".t3").Text()
		obj.SchoolType = s.Find(".t4").Text()
		obj.StarLevel = s.Find(".t5").Text()
		obj.EnrollOrder = "本科第一批"
		array = append(array, obj)

	})

	return array
}


func main() {
	// not finished, missing elements
}
