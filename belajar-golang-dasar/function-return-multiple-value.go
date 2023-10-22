package main

import "fmt"

func getFullName() (string, string) {
	return "Eko", "Khannedy"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	//ignore variable multiple value
	name, _ := getFullName()
	fmt.Println(name)
}
