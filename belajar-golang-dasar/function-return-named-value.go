package main

import "fmt"

func getFullname2() (firstName string, middleName string, lastName string) {
	firstName = "Eko"
	middleName = "Kurniawan"
	lastName = "Khannedy"
	return
}

func main() {
	a, b, c := getFullname2()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
