package main

import (
	"fmt"
	"sort"
)


func main(){
	fmt.Println("Hello World")

	var fruits=[]string{"apple", "banana", "orange","grape"}
	fmt.Printf("type of fruits is %T\n", fruits)
	fmt.Println("Length of fruits array is: ", len(fruits))

	fruits= append(fruits, "mango","papaya")
	fmt.Println("Fruits array is: ", fruits)

	// fruits= append(fruits[1:3],fruits[4:]...)
	fruits= append(fruits[1:3])
	fmt.Println("Fruits array is: ", fruits)


	highscores:= make([]int,4)

	highscores[0]= 234
	highscores[1]= 945
	highscores[2]= 465
	highscores[3]= 867
	highscores= append(highscores, 555,444,343)

	fmt.Println("Highscores array is: ", highscores)
	fmt.Println("Length of highscores array is: ", len(highscores))

	sort.Ints(highscores)
	fmt.Println("Sorted Highscores array is: ", highscores)


	//how to remove a value from slice based on index
	var courses=[]string{"maths", "science", "history","geography","english","hindi"}
	fmt.Println("Courses array is: ", courses)

	var index int=2
	courses=append(courses[:index], courses[index+1:]...)

	fmt.Println("Courses array is: ", courses)

}