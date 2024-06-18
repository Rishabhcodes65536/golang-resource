package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)


func main() {

   fmt.Println("Hello, World!")
	PerformGetRequest()
}

func PerformGetRequest(){
	const myurl="http://localhost:8000/get"
	response,err:=http.Get(myurl)
	checkNilErr(err)
    defer response.Body.Close()
	fmt.Println("Status of Get request:",response.StatusCode)
	fmt.Println("Content length is", response.ContentLength)

	var responseString strings.Builder
	content,_ :=io.ReadAll(response.Body)
	byteCount,_ := responseString.Write(content)
	// fmt.Println(content)
	fmt.Println("Byte count is", byteCount)
	fmt.Println(responseString.String())
	fmt.Println(string(content))
}
func checkNilErr(err error){
if err!=nil{
   panic(err)
}
}