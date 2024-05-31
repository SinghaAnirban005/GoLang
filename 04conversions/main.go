package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for our Pizza")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	//Handling errors....
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to you rating", numRating+1)
	}

}
