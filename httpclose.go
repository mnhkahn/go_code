package main

import "fmt"
import "net/http"
import "io/ioutil"

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://www.cyeam.com", nil)

	// NOTE this !!
    // Close indicates whether to close the connection after
    // replying to this request (for servers) or after sending
    // the request (for clients).
	req.Close = true

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("user", "pass")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}else {
        fmt.Println(string(response))
    }

}
