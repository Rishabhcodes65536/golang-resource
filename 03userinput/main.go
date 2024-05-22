package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Print out the rating for our pizza")

	// comma ok || error ok
	rating, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", rating)
	fmt.Printf("Type of rating is %T ", rating)
}
