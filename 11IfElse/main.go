package main

import "fmt"

func main() {
	loginCount := 25

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else {
		result = "Exacty 10 login count"
	}

	fmt.Println(result)

}
