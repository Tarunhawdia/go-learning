package main

import "fmt"

func main() {
	fmt.Println("hello from functions")
	greeeter()

	result := adder(3, 5)

	fmt.Println("result is", result)

	ans, myMessage := proAdder(8, 2, 3, 4, 5)
	fmt.Println("proAdder result is", ans)
	fmt.Println("my message is ", myMessage)

}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "here is the total"
}

func greeeter() {
	fmt.Println("namaste from golang")
}
