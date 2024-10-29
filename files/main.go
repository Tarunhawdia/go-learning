package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("hello	from files")

	content := "this is a good content needs to be in the file"

	file, err := os.Create("./myfiles.txt")
	checkNilErr(err)

	lenght, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("lenght of the file is ", lenght)
	file.Close()

	readFile("./myfiles.txt")

}

func readFile(filename string) {

	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("text data inside file is\n", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
