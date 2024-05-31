package main

import "fmt"

// creating a constant
const LoginToken string = "dhgfksd" //Public

func main() {
	var username string = "Anirban"
	fmt.Println(username)
	fmt.Printf("The variable is of type : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("The variable is of type %T \n", isLoggedIn)

	var smallFloat float64 = 255.233443472934729
	fmt.Println((smallFloat))
	fmt.Printf("The variable id of type: %T \n", smallFloat)

	//default values and some aliasses..
	//By default int will have 0 as value....
	var variable int
	fmt.Println(variable)
	fmt.Printf("The variable id of type: %T \n", variable)

	//implicit type..
	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style
	numberofUser := 3000
	fmt.Println(numberofUser)

}
