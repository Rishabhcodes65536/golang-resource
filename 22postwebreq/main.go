package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)


func main() {

   fmt.Println("Hello, World!")
   
	PerformPostFormRequest()
}


func checkNilErr(err error){
if err!=nil{
   panic(err)
}
}

func PerformPostJSONRequest(){
	const myurl="http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
	{
		"coursename":"golang",
		"price":0,
		"platform":"leetcode.com"
	}
	`)

	response,err := http.Post(myurl,"application/json",requestBody)
	checkNilErr(err)

	defer response.Body.Close()
	io.ReadAll(response.Body)
}

func PerformPostFormRequest(){
	const myurl="http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname","Rishabh")
	data.Add("lastname","Jain")
	data.Add("Email","Rishabh@go,dev")

	response,err := http.PostForm(myurl,data)

	checkNilErr(err)
	defer response.Body.Close()
	content,_ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}