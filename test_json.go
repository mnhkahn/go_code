package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	jsonstr := `{"10000000000":10000000000,"111":1}`
	res := make(map[string]interface{})
	if err := json.Unmarshal([]byte(jsonstr), &res); err != nil {
		panic(err)
	}
	for k, v := range res {
		fmt.Println(k, v, reflect.TypeOf(v))
	}

	// http://stackoverflow.com/questions/16946306/preserve-int64-values-when-parsing-json-in-go
	/*
		bool, for JSON booleans
		float64, for JSON numbers
		string, for JSON strings
		[]interface{}, for JSON arrays
		map[string]interface{}, for JSON objects
		nil for JSON null
	*/
	d := json.NewDecoder(bytes.NewBuffer([]byte(jsonstr)))
	d.UseNumber()
	if err := d.Decode(&res); err != nil {
		panic(err)
	}
	for k, v := range res {
		fmt.Println(k, v, reflect.TypeOf(v)) // json.Number
	}
}
