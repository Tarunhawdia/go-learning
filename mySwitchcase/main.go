package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")
	var num int
	fmt.Scanln(&num)

	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Unknown Number")
	}
}
