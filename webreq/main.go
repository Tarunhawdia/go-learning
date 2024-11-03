package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("hello from webverb")
	performGetRequest()
}

func performGetRequest() {
	const myurl = "http://loaclhost:3000/get"

	resp, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println("status code", resp.StatusCode)
	fmt.Println("content lenght", resp.ContentLength)

	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(content))

}
