package main

import "fmt"
import "reflect"

type Cyeam struct {
	Url   string `json:"url"`
	Other string `json:"-"`
}

func main() {
	c := Cyeam{Url: "blog.cyeam.com", Other: "..."}
	var t reflect.Type
	t = reflect.TypeOf(c)
	var v reflect.Value
	v = reflect.ValueOf(c)
	json := "{"
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("json") != "-" {
			json += "\"" + t.Field(i).Tag.Get("json") + "\":\"" + v.FieldByName(t.Field(i).Name).String() + "\""
		}
	}
	json += "}"
	fmt.Println(json)
}
