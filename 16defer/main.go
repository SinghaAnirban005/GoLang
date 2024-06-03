package main

import "fmt"

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")

	myDefer()

}

// Based on LIFO principle (Stack) what defer does is it puts the line of code just before the end of block of code

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
