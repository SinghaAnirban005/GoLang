package main

import "fmt"

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")

	myDefer()

}

// Based on LIFO principle

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
