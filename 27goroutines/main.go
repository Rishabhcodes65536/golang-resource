package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"test"}

var wg sync.WaitGroup  //pointer

var mut sync.Mutex

func main() {

//    go greeter("hello")
//    greeter("world")

	websiteList := []string{
		"https://google.com",
		"https://linkedin.com",
		"https://go.dev",
		"https://x.com",
	}

	for _,websites := range websiteList{
		go getHttpResponse(websites)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}


func greeter(s string){
	for i:=0;i<5;i++{
		time.Sleep(3*time.Millisecond)
		fmt.Println(s)
	}
}

func checkNilErr(err error){
	if err!=nil{
	panic(err)
	}
}

func getHttpResponse(url string) {
	defer wg.Done()

	res,err := http.Get(url)
	checkNilErr(err)
	mut.Lock()
	signals=append(signals,url)
	mut.Unlock()
	fmt.Printf("%d status code for %s\n",res.StatusCode,url)
}