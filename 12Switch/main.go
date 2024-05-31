package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")

	case 2:
		fmt.Println("You can move 2 spot")

	case 3:
		fmt.Println("ghjgjgj")

	case 4:
		fmt.Println("fjshdkfjsdbgjf")
	case 5:
		fmt.Println("HDFKDF")

	case 6:
		fmt.Println("fjdf")

	default:
		fmt.Println("fkdhfjkdh")
	}

}
