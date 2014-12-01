package main

import (
	// "bytes"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"os/exec"
	// "strings"
	// "time"
)

const bingURL = `http://cn.bing.com`
const file_path = `/tmp/wallpaper.jpg`

func main() {
	log := logs.NewLogger(10000)
	log.SetLogger("file", `{"filename":"/tmp/wallpaper.log"}`)
	log.Info("Start===================")
	v := Bing{}
	req := httplib.Get("http://www.bing.com/HPImageArchive.aspx?format=json&idx=0&n=1")
	err := req.ToXml(&v)

	if len(v.Images) > 0 && err == nil {
		pic := httplib.Get(bingURL + v.Images[0].Url)
		fmt.Println(file_path)
		err = pic.ToFile(file_path)
		if err != nil {
			log.Error("save file error %v.............", err)
		} else {
			log.Info("download wallpeper success..........")
			cmd := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", fmt.Sprintf("file:%s", file_path))
			err := cmd.Start()
			log.Info("Waiting for command to finish.........")
			err = cmd.Wait()
			log.Info("Command finished with error: %v", err)
			if err != nil {
				log.Error("%v............", err)
			}
			log.Info("set wallpeper success..........")
		}
	}
	log.Info("End===================")
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
