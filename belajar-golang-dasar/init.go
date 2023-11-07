package main

import (
	"fmt"
	"pemrograman-go-lang-pemula-sampai-mahir-udemy-programmer-zaman-now/belajar-golang-dasar/database"
	_ "pemrograman-go-lang-pemula-sampai-mahir-udemy-programmer-zaman-now/belajar-golang-dasar/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
