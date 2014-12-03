package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"regexp"
)

func main() {
	req := httplib.Get("http://liyinghuan.baijia.baidu.com/article/2011")
	res, err := req.String()
	if err != nil {
		panic(err)
	}

	title := ""
	re_h1 := regexp.MustCompile(`<h1(.*?)[^>]*?>.*?<\/h1>`)
	h1s := re_h1.FindAllString(res, -1)
	for _, h1 := range h1s {
		if len(h1) > 9 {
			title = RemoveHtml(h1)
			break
		}

	}

	fmt.Println(title)

	re_html := regexp.MustCompile("(?is)<.*?>")
	blank := re_html.ReplaceAllString(res, "")
	fmt.Println(blank)
}

func RemoveHtml(src string) string {
	re_html := regexp.MustCompile("(?is)<.*?>")
	return re_html.ReplaceAllString(src, "")
}
