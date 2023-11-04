package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}
	address2 := address1

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)

	address3 := Address{
		City:     "Jakarta",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}

	//pointer
	address4 := &address3

	address4.City = "Cimahi"

	fmt.Println(address3)
	fmt.Println(address4)

	///

	address5 := Address{
		City:     "Jakarta",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}

	address6 := &address5

	address6.City = "Cimahi"
	address6 = &Address{
		City:     "Malang",
		Province: "Jawa Timur",
		Country:  "Indonesia",
	}

	fmt.Println(address5)
	fmt.Println(address6)

	///
	address7 := Address{
		City:     "Jakarta",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}

	address8 := &address7
	address9 := &address7

	address8.City = "Cimahi"
	*address8 = Address{
		City:     "Malang",
		Province: "Jawa Timur",
		Country:  "Indonesia",
	}

	fmt.Println(address7)
	fmt.Println(address8)
	fmt.Println(address9)

	var address10 *Address = new(Address)
	address10.City = "Bandung"
	fmt.Println(address10)
}
