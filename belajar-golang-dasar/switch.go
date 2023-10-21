package main

import "fmt"

func main() {

	name := "Joko"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
		fmt.Println("Hello Eko")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Kenalan dong")
		fmt.Println("Kenalan dong")

	}

	//switch dengan short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")

	}

	//switch tanpa kondisi, mirip dengan if else
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
