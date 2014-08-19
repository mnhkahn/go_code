package main

import (
	"fmt"
	"github.com/mnhkahn/qrgo"
	"io/ioutil"
)

func main() {
	c, err := qr.Encode("lichao", qr.L)
	pngdat := c.PNG()

	if err == nil {
		ioutil.WriteFile("qr.png", pngdat, 0666)
	} else {
		fmt.Println(err)
	}
}
