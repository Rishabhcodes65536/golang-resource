package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount >10 {
		result = "Watch out"
	} else{
		result = "Exact 10 count"
	}

	fmt.Println(result)

	if num:=3;num<10 {
		fmt.Println("Num less than 10")
	} else {
		fmt.Println("Num greater than or equal to 10")
	}
}
