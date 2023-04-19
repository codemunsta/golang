package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("http://localhost")
	if err == nil {
		all, err := ioutil.ReadAll(response.Body)
		if err == nil {
			fmt.Println(string(all))
		}
	}
}
