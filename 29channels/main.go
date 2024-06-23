package main

import (
	"fmt"
	"sync"
)


func main() {

   fmt.Println("Channels in Golang")

   myCh := make(chan int, 1)  //1 meaning the buffered channel
   wg := &sync.WaitGroup{}


	//ill allow to pass on the value if someone is listening fatal error: all goroutines are asleep - deadlock! CLASSIC Error
	//i cannot pass the value coz guy listening to that is in line below i.e. 21th line

//    myCh <- 5
//    fmt.Println(<-myCh)
	wg.Add(2)
	//Receive only
	go func (ch <-chan int, wg *sync.WaitGroup){
		// close(myCh)
		val, isChanelOpen := <-myCh

		fmt.Println(isChanelOpen)
		fmt.Println(val)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh,wg)

	//send only channel
	go func (ch chan<- int, wg *sync.WaitGroup){
		// myCh <- 0
		close(myCh)
		// myCh <- 0
		// myCh <- 6
		wg.Done()
	}(myCh,wg)

	wg.Wait()

}


func checkNilErr(err error){
if err!=nil{
   panic(err)
}
}