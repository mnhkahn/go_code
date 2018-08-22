// Package jsonnumber
package main

import (
	"encoding/json"
	"log"
)

func main() {
	var raw = []byte(`{"a":1}`)
	var raw_string = []byte(`{"a":"2"}`)

	s1 := new(S1)
	err := json.Unmarshal(raw, s1)
	log.Println(err, s1.A)

	s2 := new(S2)
	err = json.Unmarshal(raw_string, s2)
	log.Println(err, s2.A)

	s21 := new(S2)
	err = json.Unmarshal(raw, s21)
	log.Println(err, s21.A)

	s3 := new(S3)
	err = json.Unmarshal(raw, s3)
	log.Println(err, s3.A)

	s31 := new(S3)
	err = json.Unmarshal(raw_string, s31)
	log.Println(err, s31.A)
}

type S1 struct {
	A int `json:"a"`
}

type S2 struct {
	A int `json:"a,string"`
}

type S3 struct {
	A json.Number `json:"a"`
}
