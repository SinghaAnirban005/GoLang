package main

import "fmt"

func main() {
	// No inheritance in golang..., No super or parent

	anirban := User{"Anirban", "anirban@gamil.com", true, 12}
	fmt.Println(anirban)

	fmt.Printf("anirban's details are: %+v\n", anirban)

	fmt.Printf("Name of user is %v\n", anirban.Name)

	Anirban := User{"Singha", "anirban.dev", true, 18}

	fmt.Printf("\nThe usage status of user %v is %v\n", Anirban.Name, Anirban.Status)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getStatus() {
	fmt.Println("The status of user is", u.Status)
}
