package main

import "fmt"

// jwtToken :="mySecretToken"  This wont work

const LoginToken="rjndunkufbnu"
//capital first letter 'L' has a significance as it is alisased to public

func main() {
	var name string = "Hello ji"
	fmt.Println(name)
	fmt.Printf("Variable is of type: %T \n",name)

	var isLogged bool = true
	fmt.Println(isLogged)
	fmt.Printf("Variable is of type: %T \n",isLogged)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n",smallVal)
	
	var smallFloat float32 = 255.3489328903289
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n",smallFloat)
	
	//default values and aliases

	var anotherVariable int;
	fmt.Println(anotherVariable);
	fmt.Printf("Variable is of type: %T \n",anotherVariable)

	//implicit type
	var website="www.knolly.co.uk"
	fmt.Println(website)

	//no var style
	//inside any method it is allowed to use walrus operator :=
	numberOfUser := 1000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
}
