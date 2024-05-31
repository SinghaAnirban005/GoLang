package main

import (
	"fmt"
	"time"
)

// In order to build this application as .exe --> IN TERMINAL --> GOOS="windows" go build

func main() {
	fmt.Println("Welcome to time-study")

	presentTime := time.Now()
	fmt.Println(presentTime)
	// to format the time we must remember the following values...
	fmt.Println(time.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}

//Memory Managaement :
// Using new() --> Allocates memory but does not initialize and make() --> allocates memory and initializes it
// Garbage Collection happens automatically
// Memory management also happens automatically...
