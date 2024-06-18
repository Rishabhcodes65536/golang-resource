package main

import (
	"encoding/json"
	"fmt"
)


type course struct{
	Name string 	`json:"coursename"`
	Price int
	Platform string		`json:"website"`
	Password string		`json:"-"`
	Tags []string		`json:"tags,omitempty"`
}

func main() {

   fmt.Println("Welcome to JSON video")

	// EncodeJSON()
	DecodeJSON()
}


func EncodeJSON(){
	lcoCourses := []course{
		{"ReactJS Bootcamp",299,"LearnCodeOnline.in","abc123",[]string{"webdev","js"}},
		{"ExpressJS Bootcamp",199,"LearnCodeOnline.in","bcd123",[]string{"Backend","js"}},
		{"MERNJS Bootcamp",599,"LearnCodeOnline.in","efg123",[]string{"Full stack","database"}},
		{"AngularJS Bootcamp",99,"LearnCodeOnline.in","ef4hefwkg123",nil},
	}

	finalJSON,err := json.MarshalIndent(lcoCourses, "","\t")

	checkNilErr(err)

	fmt.Printf("%s\n", finalJSON)


}

func DecodeJSON(){
	jsonDataFromWeb := []byte(`
	{
    "coursename": "ReactJS Bootcamp",
    "Price": 299,
    "website": "LearnCodeOnline.in",
    "tags": ["webdev","js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid{
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	}


	//some cases where you just want to add data to key value pair

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k,v := range myOnlineData{
		fmt.Printf("Key is %v and value is %v and type is  %T\n",k,v,v)
	}
}

func checkNilErr(err error){
if err!=nil{
   panic(err)
}
}