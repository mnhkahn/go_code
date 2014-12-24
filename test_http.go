package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	// "github.com/franela/goreq"
	// "io/ioutil"
	// "net/http"
)

func main() {
	// resp, _ := http.Get("http://localhost:8080/gzip")
	// defer resp.Body.Close()

	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	req := httplib.Get("http://192.168.3.151:8180/solr/judeals/select?q=*%3A*&wt=json&indent=true")
	body1, _ := req.String()
	fmt.Println(body1)

	// res, _ := goreq.Request{
	// 	Method:      "GET",
	// 	Uri:         "http://localhost:8080/gzip",
	// 	Compression: goreq.Gzip(),
	// }.Do()
	// body2, _ := res.Body.ToString()
	// fmt.Println(body2)

	// res1, _ := goreq.Request{
	// 	Method: "GET",
	// 	Uri:    "http://192.168.3.151:8180/solr/judeals/select?q=*%3A*&wt=json&indent=true",
	// 	// Compression: goreq.Gzip(),
	// }.Do()
	// body3, _ := res1.Body.ToString()
	// fmt.Println(body3)
}
