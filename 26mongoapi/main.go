package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Rishabhcodes65536/mongoapi/router"
)

func main(){
	fmt.Println("MongoDB api")
	r:=router.Router()
	fmt.Println("Server is getting started...")

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000")

}