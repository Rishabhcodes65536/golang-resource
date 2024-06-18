package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Course struct {
CourseId string `json:"courseid"`
CourseName string `json:"courseid"`
CoursePrice int `json:"-"`
Author *Author  `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

var courses []Course

//middleware, helper -file

func (c *Course) isEmpty() bool{
	// return c.CourseId =="" && c.CourseName==""
	return c.CourseId ==""
}

func main() {

   fmt.Println("You are on the API initialization")

   r:=mux.NewRouter()

   //seeding	
   courses = append(courses, Course{CourseId: "2",CourseName: "Reactjs",CoursePrice: 699,Author:&Author{Fullname: "Rishabh Jain",Website: "knolly.ai"}})
   courses = append(courses, Course{CourseId: "4",CourseName: "MERN",CoursePrice: 199,Author:&Author{Fullname: "Rishabh Jain",Website: "go.dev"}})

   //routing
   r.HandleFunc("/",serveHome).Methods("GET")
   r.HandleFunc("/courses",getAllCourses).Methods("GET")
   r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
   r.HandleFunc("/course",createOneCourse).Methods("POST")
   r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
   r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE")


   //listen to port
   log.Fatal(http.ListenAndServe(":4000",r))


}


func checkNilErr(err error){
if err!=nil{
   panic(err)
}
}


//contollers

//serve home file

//for good practice just name as course request would handle the case

func serveHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Welcome to Rishabh's Go server</h1>"))
}

func getAllCourses(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get one courses")
	w.Header().Set("Content-type","application/json")

	// grab id from uses

	params := mux.Vars(r)
	fmt.Println(params)

	//loop through our fake db courses to find matching id and return response

	for _,course := range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}


func createOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Create one courses")
	w.Header().Set("Content-type","application/json")

	//what if body is empty
	if r.Body == nil{
		json.NewEncoder(w).Encode("Please send some data")	
		return
	}

	// what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty(){
		json.NewEncoder(w).Encode("Invalid request")
		return
	}
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses,course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Update one courses")
	w.Header().Set("Content-type","application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop id,remove,add with my ID
	for index,course :=range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index],courses[index+1:]...)
			var newCourse Course
			_ = json.NewDecoder(r.Body).Decode(&newCourse)
			newCourse.CourseId = params["id"]
			courses=append(courses, newCourse)
			json.NewEncoder(w).Encode(newCourse)
			return
		}
	}

	//send a response if id not found
	json.NewEncoder(w).Encode("No course found with given id")
}



func deleteOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Delete one courses")
	w.Header().Set("Content-type","application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop id,remove,add with my ID
	for index,course :=range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index],courses[index+1:]...)
			break
		}
	}
}