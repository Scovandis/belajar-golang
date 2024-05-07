package main

import (
	"fmt"
)

func main() {
	fmt.Println(fullsName())

	firstName, _, lastName := fullNames()

	fmt.Println(firstName)
	fmt.Println(lastName)
}

func fullsName() (string, string) {
	return "andre", "sih"
}

func fullNames() (string, string, string) {
	return "andre", "middel", "Lastname"
}
