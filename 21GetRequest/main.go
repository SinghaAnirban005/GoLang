package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web-verb video...")

	performGetRequest()

	PerformPostJsonRequest()

}

func performGetRequest() {
	const myURL = "http://localhost:8000/get"

	response, err := http.Get(myURL)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content length is :", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

	//Another Method..
	var responseString strings.Builder
	newcontent, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(newcontent)

	fmt.Println("Byte Count is ", byteCount)
	fmt.Println(responseString.String())

}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
	{
		"coursename": "Golang",
		"price" : "2500",
		"user" : "anirban",
	}`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
