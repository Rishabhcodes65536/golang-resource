package main

import "fmt"

func main() {
	// var ptr *int
	// fmt.Println("Value of ptr is: ", ptr)

	myNum := 50

	var ptr = &myNum

	fmt.Println("Value of actual pointer is: ", ptr);
	fmt.Println("Value of actual pointer is: ", *ptr);

	*ptr = *ptr *2
	fmt.Println("New value is: ",myNum)
}
