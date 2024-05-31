package main

import (
	"fmt"
	"sort"
)

func main() {

	var vegList = []string{"Apple", "Mango", "Tomato"}

	fmt.Printf("The type of vegList is %T\n", vegList)

	vegList = append(vegList, "Paneer", "Milk")

	fmt.Println(vegList)
	//Slicing up the slice
	vegList = append(vegList[1:])
	fmt.Println(vegList)

	vegList = append(vegList[1:3])
	fmt.Println(vegList)

	//Another way of declarung slices
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 123
	highScores[2] = 345
	highScores[3] = 678

	//Below using the append func will reallocate memory and increase size

	highScores = append(highScores, 9, 12, 23)
	fmt.Println(highScores)

	//sorting slices
	sort.Ints(highScores)
	fmt.Println(highScores)

	// how to remove values from slices based on index
	var courses = []string{"reactjs", "js", "vuejs", "swift", "python"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
