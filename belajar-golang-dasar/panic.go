package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
}

func runaApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}

	fmt.Println("Aplikasi berjalan")
}
func main() {
	runaApp(true)
	fmt.Println("Eko")
}
