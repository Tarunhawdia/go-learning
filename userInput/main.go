package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
		
	welcome := "welcome to the user input program"	
	fmt.Println(welcome)

	redear:= bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating")

	//comma ok syntex || error ok

	input, _:= redear.ReadString('\n')
	fmt.Printf("thanks for rating is: %T", input);

}
