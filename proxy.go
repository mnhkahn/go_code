package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/franela/goreq"
	"log"
	"time"
)

var proxys []string

func main() {
	for _, proxy := range proxys {
		res, err := goreq.Request{
			Uri:       "http://www.baidu.com/",
			UserAgent: "Cyeambot",
			Proxy:     proxy,
			Timeout:   5 * time.Second,
		}.Do()
		goreq.SetConnectTimeout(5 * time.Second)
		if err != nil || res.Body == nil {
			log.Printf("%s fail.\n", proxy)
		} else {
			log.Printf("%s success.\n", proxy)
		}
	}
}

func InitProxys() {
	doc, err := goquery.NewDocument("http://proxy.com.ru/list_1.html")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("body > center > font > table:nth-child(1) > tbody > tr > td:nth-child(2) > font > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
		if i > 0 {
			s.Find("td").Each(func(j int, s *goquery.Selection) {
				if j == 1 {
					proxys = append(proxys, "http://"+s.Text())
				} else if j == 2 {
					proxys[i-1] = proxys[i-1] + ":" + s.Text()
				}
			})
		}
	})
}

func init() {
	InitProxys()
}
