package main

import "fmt"

func main() {

	//Using Pointers
	myNumber := 23
	var ptr = &myNumber

	fmt.Println("The value of ptr is ", ptr)
	fmt.Println("The value of pointer is", *ptr)

	// Opearations
	*ptr = *ptr + 2
	fmt.Println("The value of ptr after ops is ", *ptr)

}
