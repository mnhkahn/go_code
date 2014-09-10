package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("RemoteAddr: ", r.RemoteAddr)
		fmt.Println("Params: ", r.URL.Query().Get("param0"))
		fmt.Println("Method: ", r.Method)
		fmt.Println("Cookie: ", r.Cookies())
		fmt.Println("UserAgent: ", r.UserAgent())
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Println("Header: ", r.Header.Get("header0"))
		fmt.Println("Body: ", string(body))

		io.WriteString(w, "hello, world!\n")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
