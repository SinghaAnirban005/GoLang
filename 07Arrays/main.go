package main

import "fmt"

func main() {

	var array [4]string

	array[0] = "Tomato"
	array[1] = "Cake"
	array[2] = "Potato"

	fmt.Println(array)
	//Length should have been 3 but this len func will return the actual length of array..
	fmt.Println("The length of the array is ", len(array))

	//another way of declaring and assigning values 
	var vegList = [5]string{"Paneer", "Tomato", "Capsicum"}
	fmt.Println(vegList)

}
