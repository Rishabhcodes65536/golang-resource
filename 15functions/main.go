package main

import "fmt"

func main() {
	fmt.Println("This is function tutorial")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is", result)

	proRes,mssg := proAdder(2,3,4,5,6,6,7,8)

	fmt.Println("Pro result is", proRes)
	fmt.Println("Pro message is", mssg)
}
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int,string) {
	total :=0
	for _,val := range values{
		total+=val
	}
	return total,"Hello pro Adder user"
}

func greeter() {
	fmt.Println("This is the called function")

}
