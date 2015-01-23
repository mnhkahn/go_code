package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
)

var proxys []string

func main() {
	log.Println(proxys)
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
