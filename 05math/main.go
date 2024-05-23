package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	"crypto/rand"
	// "time"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	var inp1 int = 10
	var inp2 float64 = 4.5
	fmt.Println("The sum is: ", inp1+int(inp2))

	//random number
	// rand.Seed(time.Now().UnixNano()) //i didnt require this it is golang 1.22.3 v
	// fmt.Println(rand.Intn(5))

	//random from crypto

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
