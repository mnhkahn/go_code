package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
)

func main() {
	j := simplejson.New()
	j.Set("user", 1)
	res, err := j.MarshalJSON()
	fmt.Println(string(res), err)
}
