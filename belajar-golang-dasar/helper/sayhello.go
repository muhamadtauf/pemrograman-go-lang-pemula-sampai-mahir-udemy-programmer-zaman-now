package helper

import "fmt"

// diawali dengan huruf kecil, hanya bisa diakses di package yang sama
var version = "1.0.0"

// diawali dengan huruf besar, bisa diakses dari package lain
var Application = "Golang"

// diawali dengan huruf kecil, hanya bisa diakses di package yang sama
func sayGoodBye(name string) string {
	return "Good Bye " + name
}

// diawali dengan huruf besar, bisa diakses dari package lain
func SayHello(name string) {
	fmt.Println("Hello", name)
}

func Contoh() {
	sayGoodBye("Eko")
	fmt.Println(version)
}
