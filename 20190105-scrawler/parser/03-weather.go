package main

import (
	"../util/filer"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	html := filer.Readfile("bom.html")

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	// get all under divs
	//dom.Find("body").Each(func(i int, selection *goquery.Selection) {
	//	fmt.Println(selection.Text())
	//
	//	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	//	fmt.Println(match)
	//})

	dom.Find("script").Remove()

	//dom.Find("body").Each(func(i int, selection *goquery.Selection) {
	//fmt.Println(selection.Text())

	//r, _ := regexp.Compile("(.*)\n\n(.*)Now[0-9.]*°\n\n(.*)\n\n(.*)\n\n(.*)")
	//fmt.Println(r.FindAllString(selection.Text(), -1))

	//})

	txt := dom.Find("body").Text()
	//fmt.Println(txt)
	r, _ := regexp.Compile("(.*)\n\n(.*)Now[0-9.]*°\n\n(.*)\n\n(.*)\n\n(.*)")
	units := r.FindAllString(txt, -1)
	//fmt.Println(units)

	result := make(map[string][]string)

	rcity, _ := regexp.Compile("(.*)")
	rtemp, _ := regexp.Compile("(?:Now)([0-9.]*)°")
	rwind, _ := regexp.Compile("([A-Z]*) ([0-9]*)(?:km/h)")
	rmax, _ := regexp.Compile("(?:Max)(.*)([0-9]*)°")
	rminMax, _ := regexp.Compile("([0-9]{2})°(?:.*)([0-9]{2})")
	for _, weather := range units {
		fmt.Println(weather)
		fmt.Println("------------------------")
		city := rcity.FindAllString(weather, -1)[0]
		city = strings.ToLower(city)
		temperature := rtemp.FindStringSubmatch(weather)[1]
		wind := rwind.FindStringSubmatch(weather)
		direction := wind[1]
		//item["speed"] = wind[2]
		Dspeed, _ := strconv.ParseFloat(wind[2], 64)
		speed := strconv.FormatFloat(Dspeed/1.852, 'f', 1, 64) + " knots"
		max := rmax.FindStringSubmatch(weather)
		iMin := "null"
		iMax := "null"
		if len(max) < 1 {
			max := rminMax.FindStringSubmatch(weather)
			iMin = max[1]
			iMax = max[2]
		}else {
			iMax = max[1]
		}
		
		ss := strings.Split(weather, "\n")
		desc := ss[(len(ss)-1)]
		result[city] = []string{city, temperature, direction, speed, iMin, iMax, desc}
	}

	fmt.Println(result["hobart"])

}
