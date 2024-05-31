package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("File Handling in Golang...")

	content := "This needs to go in a file--> Learn Code Online.in"
	// to create a file
	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err)
	}
	// to write the content into file
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("The length is: ", length)
	defer file.Close()

	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	// To read the contents of the file into a variable..
	databyte, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("text data inside file is :", string(databyte))
}
