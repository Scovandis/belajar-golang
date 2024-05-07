package main

import "fmt"

func main() {

	a, b, _ := returnNamed()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(returnNamed())
}

func returnNamed() (firstName, middelName, lastName string) {
	firstName = "Andre"
	middelName = "Middle Na"
	lastName = "Last"

	return
}
