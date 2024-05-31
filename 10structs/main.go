package main

import "fmt"

func main() {
	// No inheritance in golang..., No super or parent

	anirban := User{"Anirban", "anirban@gamil.com", true, 12}
	fmt.Println(anirban)

	fmt.Printf("anirban's details are: %+v", anirban)

	fmt.Printf("Name of user is %v", anirban.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
