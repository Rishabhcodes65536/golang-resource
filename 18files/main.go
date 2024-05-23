package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in GoLang")

	content := "This needs to go in a file - www.knolly.co.uk/offer"

	file, err := os.Create("./mygofile.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	length, err := io.WriteString(file, content)

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err);
	fmt.Println("The length of file is", length)
	fmt.Println("File name:", file.Name())
	readFile(file.Name())
}


func readFile(filename string){
	databyte,err := os.ReadFile(filename)

	// if err!=nil{
	// 	panic(err)
	// }
	checkNilErr(err)

	fmt.Println("Text data inside file is(unconverted) ",databyte);
	fmt.Println("Text data inside file is ",string(databyte));
}


func checkNilErr(err error){
	if err!=nil{
		panic(err)
	}
}