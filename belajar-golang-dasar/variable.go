package main

import "fmt"

func main() {
	var name string

	name = "Eko Kurniawan"
	fmt.Println(name)

	name = "Eko Khannedy"
	fmt.Println(name)

	//penulisan varible tidak wajib menukiskan tipe datanya apabila langsung diberi nilai
	var friendName = "budi"
	fmt.Println(friendName)

	var age = 30
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	//multiple variable
	var (
		firstName = "Eko"
		lastName  = "Khannedy"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
