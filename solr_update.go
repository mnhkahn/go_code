package main

import (
	"fmt"
	"github.com/franela/goreq"
	"time"
)

func main() {
	// Add()
	Del()
}

func Del() {
	item := DelSolr{}
	item.Del.Query = `title:cyeam.com`
	item.Del.CommitWithin = 1000
	req := goreq.Request{
		Method:      "POST",
		Uri:         "http:///solr/post/update?wt=json",
		ContentType: "application/json",
		Body:        item,
		// Proxy:       "http://127.0.0.1:8888",
	}
	res, err := req.Do()
	if err != nil {
		fmt.Println(err, "*****")
		return
	}
	str, _ := res.Body.ToString()
	fmt.Println("del success", str)

}

func Add() {
	item := AddSolr{}
	item.Add.Doc.Link = time.Now().String()
	item.Add.Doc.Title = "cyeam.com"
	item.Add.Overwrite = true
	item.Add.CommitWithin = 1000
	req := goreq.Request{
		Method:      "POST",
		Uri:         "http:///solr/post/update?wt=json",
		ContentType: "application/json",
		Body:        item,
		// Proxy:       "http://127.0.0.1:8888",
	}
	res, err := req.Do()
	if err != nil {
		fmt.Println(err, "*****")
		return
	}
	str, _ := res.Body.ToString()
	fmt.Println("add success", str)
}

type AddSolr struct {
	Add struct {
		Doc          SolrModel `json:"doc"`
		Overwrite    bool      `json:"overwrite"`
		CommitWithin int       `json:"commitWithin"`
	} `json:"add"`
}

type SolrModel struct {
	Link  string `json:"link"`
	Title string `json:"title"`
}

type DelSolr struct {
	Del struct {
		Query        string `json:"query"`
		CommitWithin int    `json:"commitWithin"`
	} `json:"delete"`
}
