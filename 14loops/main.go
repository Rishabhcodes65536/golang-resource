package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday","Tuesday","Wednesday","Saturday","Sunday"}

	fmt.Println(days)
	//no ++d
	// for d:=0;d<len(days);d++{
	// 	fmt.Println(days[d])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	for index, day := range days{
		fmt.Printf("Index is %v and value is %v\n", index,day)
	}

	rougueValue :=1

	for rougueValue < 10 {
		if rougueValue ==2{
			goto rj
		}
		if rougueValue ==5{
			break
		}
		fmt.Println("Value is: ",rougueValue)
		rougueValue++
	}

	rj:
		fmt.Println("Hello a label")
}
