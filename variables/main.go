package main

import "fmt"


func main(){
	var username string = "John Doe"
	fmt.Println(username);
	fmt.Printf("variable type: %T\n", username);

	var isLoggedIn bool=false 
	fmt.Println(isLoggedIn );
	fmt.Printf("variable type: %T\n", isLoggedIn);

	var smallval uint8=255
	fmt.Println(smallval);
	fmt.Printf("variable type: %T\n", smallval);


	var smallfloat uint8=255
	fmt.Println(smallfloat);
	fmt.Printf("variable type: %T\n", smallfloat);
}