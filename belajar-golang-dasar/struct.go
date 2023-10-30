package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// struct method
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)

}

func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Indonesia"
	eko.Age = 30

	fmt.Println(eko)
	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Cirebon",
		Age:     35,
	}

	fmt.Println(joko)

	budi := Customer{"Budi", "Jakarta", 35}

	fmt.Println(budi)

	// memanggil struct method
	eko.sayHi("Joko")

}
