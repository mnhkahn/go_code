package main

import (
	"fmt"
	"github.com/franela/goreq"
)

func main() {
	item := AddSolr{}
	item.Add.Doc.Id = "1234"
	item.Add.Doc.Title = "5678"
	item.Add.Overwrite = true
	item.Add.CommitWithin = 1000
	req := goreq.Request{
		Method:      "POST",
		Uri:         "http://128.199.131.129:8983/solr/post/update?wt=json",
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
	fmt.Println("success", str)
}

type AddSolr struct {
	Add struct {
		Doc          SolrModel `json:"doc"`
		Overwrite    bool      `json:"overwrite"`
		CommitWithin int       `json:"commitWithin"`
	} `json:"add"`
}

type SolrModel struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}
