package main

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"os/exec"
)

func main() {
	req := httplib.Get("http://blog.cyeam.com/golang/2014/11/29/golang_gzip/")
	res, err := req.Bytes()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("pup", `head meta`)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	// defer stdin.Close()
	var output bytes.Buffer
	cmd.Stdout = &output

	if err = cmd.Start(); err != nil { //Use start, not run
		fmt.Println("An error occured: ", err) //replace with logger, or anything you want
	}
	stdin.Write(res)
	stdin.Close()

	if err := cmd.Wait(); err != nil {
		panic(err)
	}
	fmt.Println(string(output.Bytes())) //for debug
}
