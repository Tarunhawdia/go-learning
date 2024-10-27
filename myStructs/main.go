package main

import "fmt"

func main() {
	fmt.Println("hello from structs")

	//there is no inheritance in Golang
	//no super or parent

	tarun := User{"tarun", "tarun@go.com", 34, true}

	fmt.Println(tarun)
	fmt.Printf("tarun details are %+v\n", tarun)
	fmt.Printf("tarun name is: %v and email is %v\n", tarun.Name, tarun.Email)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
