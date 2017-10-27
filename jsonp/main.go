package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", fooHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	u := r.URL.Query()
	callback := u.Get("callback")
	// callback = template.JSEscapeString(callback)

	res := make(map[string]interface{})
	res["domain"] = "cyeam.com"

	j, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if len(callback) > 0 {
		_, err = w.Write([]byte(callback))
		if err != nil {
			log.Println(err)
		}
		_, err = w.Write([]byte{'('})
		if err != nil {
			log.Println(err)
		}
	}

	_, err = w.Write(j)
	if err != nil {
		log.Println(err)
	}

	if len(callback) > 0 {
		_, err = w.Write([]byte{')'})
		if err != nil {
			log.Println(err)
		}
	}
}
