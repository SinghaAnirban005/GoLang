package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to userInput"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for our Pizza")
	// comma ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
	// The type of input is going to be String here..
	fmt.Printf("Type of rating is %T \n", input)

}
