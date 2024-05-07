package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hallo")
	sayWithHelo("DOG", spamFilter)
	sayWithHelo("Andre", spamFilter)
}

func sayWithHelo(name string, filter func(string) string) {
	fmt.Println("Hallo ", filter(name))
}

func spamFilter(name string) string {
	if name == "DOG" {
		return "***"
	} else {
		return name
	}
}
