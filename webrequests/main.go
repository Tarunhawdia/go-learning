package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://theuselessweb.com"

func main() {
	fmt.Println("hello form webrequest")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("response type is %T\n", response)
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(databytes))

}
