package main

import (
	"fmt"
)

func main(){
	fmt.Println("maps in go!")

	courses:=make(map[string]int)
	courses["js"]=2
	courses["xx"]=1
	courses["ks"]=44
	courses["py"]=3

	fmt.Println("courses are:",courses);

	fmt.Println("js price is ", courses["js"])

	delete(courses,"js")
	fmt.Println("courses are:",courses);
	fmt.Println("tarun price is ", courses["tarun"])


	for key,value :=range courses{
		fmt.Printf("for key %v value is %v\n",key,value)
	}
	
	
}