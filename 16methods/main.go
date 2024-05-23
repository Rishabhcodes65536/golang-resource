package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")

	//no inheritance in golang;no super or no parent
	rishabh := User{"Rishabh","rishabh@smart.dev",true,21}
	fmt.Println(rishabh)
	fmt.Printf("Rishabh details are: %+v\n", rishabh)
	fmt.Printf("Name is %v & email is %v\n", rishabh.Name,rishabh.Email)
	rishabh.GetStatus()
	rishabh.NewEmail()
	fmt.Printf("Name is %v & email is %v\n", rishabh.Name,rishabh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){
	fmt.Println("Is user active", u.Status)
}

func (u User) NewEmail(){
	u.Email = "rishabh@new.dev"
	fmt.Println("Updated Email is", u.Email)
}