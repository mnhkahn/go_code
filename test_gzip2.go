package main

import (
	// "bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := http.Client{}

	req1, err := http.NewRequest("GET", "http://blog.cyeam.com", nil)
	req1.Header.Add("Content-Encoding", `gzip`)
	req1.Header.Add("Accept-Encoding", `gzip`)
	resp1, err := client.Do(req1)
	if err != nil {
		panic(err)
	}
	defer resp1.Body.Close()
	body, err := ioutil.ReadAll(resp1.Body)
	fmt.Println("Length of uncompress: ", len(body))
	compressedReader, err := gzip.NewReader(resp1.Body)
	body1, err := ioutil.ReadAll(compressedReader)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length of gzip: ", len(body1))
}
