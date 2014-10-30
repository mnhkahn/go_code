package main

import (
	"fmt"
	"github.com/franela/goreq"
)

func main() {
	res, _ := goreq.Request{
		Method:      "GET",
		Uri:         "http://blog.cyeam.com",
		Compression: goreq.Gzip(),
		Proxy:       "http://localhost:8888",
	}.Do()
	body, _ := res.Body.ToString()
	fmt.Println(body)
}
