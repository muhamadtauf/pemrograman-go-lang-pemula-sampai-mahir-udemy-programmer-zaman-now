package main

import "fmt"

func main() {

	//constant tidak bisa diubah value nya
	const firstName string = "Eko"
	const lastName = "Khannedy"
	const value = 100

	//multiple constant
	const (
		age     = 20
		address = "Bandung"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
	fmt.Println(age)
}
