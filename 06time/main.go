package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of Golang documentation is killing me fr!!")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	//Turns out we always use this exact date for formatting womp-womp

	//TIME PACKAGE IS TOO CONFUSING!!!

	createdDate := time.Date(2050,12,10,12,32,10,12,time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}


//go env to find about go stuff

// GOOS = "windows/darwin/linux" go build
