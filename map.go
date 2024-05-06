package main

import (
	"fmt"
)

// func main() {
// 	person := map[string]string{
// 		"name":    "andre_sih",
// 		"address": "surabaya",
// 	}

// 	fmt.Println(person)
// 	fmt.Println(person["name"])
// 	fmt.Println(person["address"])

// }

func main() {
	book := make(map[string]string)

	book["title"] = "Learning go land"
	book["author"] = "andre sih"
	book["wrong"] = "false"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "wrong")
	fmt.Println(book)
	fmt.Println(len(book))
}
