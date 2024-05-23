package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices it's important and extensively used");

	var fruitList = []string{"Apple","Tomato","Peach"}
	fmt.Printf("Type of fruit list is %T \n", fruitList)

	fruitList = append(fruitList,"Apple","Mango")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 888
	// highScores[4] = 1000

	highScores = append(highScores, 555, 666,777)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	//how to remove from slices(Abstraction of array) based on index

	var courses = []string{"reactjs","ruby","javascript","swift","python"}
	fmt.Println(courses)
	var index int =2
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)
}