package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {

   fmt.Println("Hello, World!")
	greeter()
	r := mux.NewRouter()
    r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000",r))
}

//go mod init github.com/Rishabhcodes65536/mymodules

//go get -u github.com/gorilla/mux
//go build .
//go mod tidy  
//go list
//go list all
//go list -m  all
//go list -m -versions github.com/gorilla/mux

//go mod why github.com/gorilla/mux why is this module in my code/what are the dependencies?

//go mod graph

// go mod edit -go 1.22.3
// go mod edit -module renamed module

//go mod vendor --> node_modules similar

//go run main.go bring from webcache

//go run -mod=vendor main.go
func greeter(){
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Welcome to golang little one! </h1>"))
}

func checkNilErr(err error){
if err!=nil{
   panic(err)
}
}