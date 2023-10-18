package main

import "fmt"

func main() {

	//alias dari tipe data string
	type NoKtp string

	//alias dari tipe data boolean
	type Maarried bool

	//membuat variable dari alias string
	var noKtpEko NoKtp = "1232131231231"
	var marriedStatus Maarried = true

	fmt.Println(noKtpEko)
	fmt.Println(marriedStatus)
}
