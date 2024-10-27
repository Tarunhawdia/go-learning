package main

import "fmt"

func main(){
	fmt.Println("hello from myPointers")

	// var ptr *int

	// fmt.Println("Value of ptr is: ", ptr)

	myNumber:=12

	var myNumptr = &myNumber

	fmt.Println("Value of pointer is: ", myNumptr)
	fmt.Println("Value of pointer is: ", *myNumptr)

	*myNumptr*=2
	fmt.Println("new value is :", myNumber)


}