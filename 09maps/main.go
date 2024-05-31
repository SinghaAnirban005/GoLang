package main

import "fmt"

func main() {
	languages := make(map[string]string)
	//KEY VALUE PAIRS...
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)

	//ACCESING ELEMENTS USISNG KEY
	fmt.Println("accesing a element:  ", languages["JS"])

	//Deleting element...
	delete(languages, "RB")
	fmt.Println("Updated list", languages)

	//Looping thorugh maps
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}

}
