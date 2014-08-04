package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type QuoteString struct {
	QString string
}

func (q QuoteString) MarshalJSON() ([]byte, error) {
	return []byte(strconv.QuoteToASCII(q.QString)), nil
}

type ColorGroup struct {
	ID     int         `json:"id,string"`
	Name   QuoteString `json:"name"`
	Colors []string
}

func main() {
	rs := []rune("golang中文unicode编码")
	j := ""
	html := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			j += string(r)
			html += string(r)
		} else {
			j += "\\u" + strconv.FormatInt(int64(rint), 16) // json
			html += "&#" + strconv.Itoa(int(r)) + ";"       // 网页
		}
	}
	fmt.Printf("JSON: %s\n", j)
	fmt.Printf("HTML: %s\n", html)

	c := ColorGroup{}
	c.ID = 20140804
	c.Name.QString = "二零一四年八月四日"
	c.Colors = append(c.Colors, "黑色")

	result, _ := json.Marshal(c)
	fmt.Println(string(result))
}
