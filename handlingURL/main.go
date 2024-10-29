package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://courses.chaicode.com/learn"

func main() {
	fmt.Println("Hello World")
	fmt.Println(myUrl)

	//parsing
	result, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("the type of query params is %T\n", qparams)

	fmt.Println(qparams["name"])

	partOfUrl := &url.URL{
		Scheme: "https",
		Host:   "www.chaicode.com",
		Path:   "learn",
	}

	anotherURL := partOfUrl.String()

	fmt.Println(anotherURL)
}
