package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Golang..")

	// Calling a function
	greet()

	result := adder(5, 6)
	fmt.Println("Result is:", result)

	proRes, message := pro(1, 2, 3, 5)
	fmt.Println(proRes)
	fmt.Println(message)
}

func greet() {
	fmt.Println("Another function body..")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// Pro function
func pro(vals ...int) (int, string) {
	total := 0

	for _, val := range vals {
		total += val
	}

	return total, "My name is Singha"
}
