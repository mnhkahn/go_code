package main

import (
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"go_code/wallpaper/win32api"
	"time"
)

const bingURL = `http://cn.bing.com`
const file_path = `C:\Users\Public\Pictures\wallpaper.jpg`
const file_path1 = `/tmp/wallpaper.jpg`

func main() {
	v := Bing{}
	req := httplib.Get("http://www.bing.com/HPImageArchive.aspx?format=json&idx=0&n=1")
	err := req.ToXml(&v)
	fmt.Println(runtime.GOOS)
	if len(v.Images) > 0 && err == nil {
		pic := httplib.Get(bingURL + v.Images[0].Url)
		// fmt.Println(pic)
		err = pic.ToFile(file_path)
		if err != nil {
			fmt.Println(err)
		} else {
			time.Sleep(5000)
			fmt.Println(win32api.SetWallPaper(file_path))
		}
	}
}

type Bing struct {
	XMLName xml.Name `xml:"images"`
	Images  []Image  `xml:"image"`
}

type Image struct {
	XMLName       xml.Name `xml:"image"`
	Startdate     string   `xml:"startdate"`
	Fullstartdate string   `xml:"fullstartdate"`
	Url           string   `xml:"url"`
}
