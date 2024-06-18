package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=dsnniosn"

func main() {

   fmt.Println("Welcome to handling URL's in GoLang")
   
   fmt.Println(myurl)

   //parsing

   result, _ := url.Parse(myurl)
   fmt.Printf("Type of result is: %T \n", result)
   fmt.Println(result.Scheme)
   fmt.Println(result.Host)
   fmt.Println(result.Path)
   fmt.Println(result.Port())
   fmt.Println(result.RawQuery)

   qparams := result.Query()
   fmt.Printf("Type of query params is %T\n", qparams)

   fmt.Println(qparams["coursename"])

   for _,val := range qparams{
		fmt.Println("Param is: ", val)
   }

   partOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=Rishabh",
   }

   anotherUrl := partOfUrl.String()
   fmt.Println(anotherUrl)
}


func checkNilErr(err error){
if err!=nil{
   panic(err)
}
}