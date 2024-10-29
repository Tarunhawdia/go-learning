package main

import "fmt"

func main() {

	//defer works as last in first out
	defer fmt.Println("here is defer")
	defer fmt.Println("One")
	defer fmt.Println("two")

	fmt.Println("hello form defer")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
