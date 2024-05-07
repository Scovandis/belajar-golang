package main

import (
	"fmt"
)

type Filter func(string) string

func main() {
	fmt.Println("Hallo")
	sayWithHelo("DOG", spamFilter)
	sayWithHelo("Andre", spamFilter)
	fmt.Println(sayAja("Andre", spamFilter))
}

func sayWithHelo(name string, filter Filter) {
	fmt.Println("Hallo ", filter(name))
}

func sayAja(name string, filter func(string) string) string {
	return fmt.Sprintf("SAY AJA YAHHH>>> %s", filter(name))
}

func spamFilter(name string) string {
	if name == "DOG" {
		return "***"
	} else {
		return name
	}
}
