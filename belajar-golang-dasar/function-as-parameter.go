package main

import "fmt"

// function type declaration
type Filter func(string) string

func sayHellloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)

}

func sayHellloWithFilter2(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)

}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}

}

func main() {
	sayHellloWithFilter("Eko", spamFilter)
	sayHellloWithFilter("Anjing", spamFilter)
	sayHellloWithFilter2("Anjing", spamFilter)
}
