package main

import "fmt"

func main() {

	days := []string{"Sunday", "Moday", "Tuesday", "Saturday"}

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, days := range days {
		fmt.Printf("index is %v and value is %v\n", index, days[index])
	}

	//Demonstration of while loop in Golang...
	rogueValue := 1

	for rogueValue < 10 {

		fmt.Println(rogueValue)
		rogueValue++

	}

	// using goto in Golang
lco:
	fmt.Println("My name is Anirban Singha")

	val := 1

	for val < 10 {
		if val == 2 {
			goto lco

			//Upon hitting goto statement the the val resets to 1
			val++
			continue
		}
		if val == 5 {
			break
		}

		fmt.Println(rogueValue)
		rogueValue++

	}

}
