package main

import (
	"fmt"
	"time"
)


func main(){
	fmt.Println("hello from handlingTime")
	presentTime:=time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createTime:=time.Date(2021, 10, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createTime)
}