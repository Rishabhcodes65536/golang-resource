package main

import "fmt"

func main() {


	//   INSHORT NO USE NO SORT METHOD PROVIDEDD APPARENTLY SO YEEAH IT ENDS!!!
	fmt.Println("Welcome to array in GoLang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	// fruitList[2] = "Cherry"
	fruitList[3] = "Peach"

	fmt.Println("Fruit List is : ", fruitList)
	fmt.Println("Fruit List length is : ", len(fruitList))

	var vegList= [3]string{"potato","onion","cucumber"};

	fmt.Println("Veggies list is: " ,vegList)

}
