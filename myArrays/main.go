package main

import "fmt"

func main(){
	fmt.Println("hello from myArrays")

	var fruits [4]string
	fruits[0]="apple"
	fruits[1]="banana"
	fruits[2]="orange"

	fmt.Println("Fruits array is: ", fruits);

	fmt.Println("Length of fruits array is: ", len(fruits))
	
	var veglist = [3]string{"potato", "tomato", "onion"}
	fmt.Println("Veglist array is: ", veglist)
	fmt.Println("Length of veglist array is: ", len(veglist))
}