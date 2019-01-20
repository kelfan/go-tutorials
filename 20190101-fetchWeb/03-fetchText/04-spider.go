package main

// http://www.voidcn.com/article/p-ogyhoiyt-bkg.html

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

const tmail_crt = "/xxxx/TMAIL.cer"

func main() {
	//tsimplehttp("https://example.com/")
	//tmailSpider()
	tOtherSpider()
}

func tmailSpider() {
	url := "https://weibots.tmall.com/category-1284634809.htm?spm=a220o.1000855.0.0.k9Wpql&search=y&catName=%D0%C2%CA%E9%B5%BD%B5%EA"
	//url = "https://www.taobao.com/"
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	client := http.Client{Transport: transport}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("get faild")
	}
	fmt.Println(resp)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func tOtherSpider() {
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", "https://weibots.tmall.com/category-1284634809.htm", nil)

	reqest.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6,zh-TW;q=0.4")
	reqest.Header.Set("Cache-Control", "max-age=0")
	reqest.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.98 Safari/537.36")
	reqest.Header.Set("Connection", "keep-alive")

	response, _ := client.Do(reqest)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
	} else {
		fmt.Println(response)
	}
}

func tsimplehttp(url string) {
	response, _ := http.Get(url)
	fmt.Println(response)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}