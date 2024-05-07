package main

import (
	"fmt"
)

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) string {
	if blacklist(name) {
		return fmt.Sprintf("Kamu di blacklist dari apps ini %s, hubungi developer !!", name)
	} else {
		return fmt.Sprintf("Selamat datang %s!! semoga harimu baik", name)
	}
}

func userRegister(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Your Blocked", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}
func main() {
	fmt.Println(registerUser("andre", validation))

	blacklist := func(name string) bool {
		return name == "anjing"
		// return name == "DOG",
		// return name == "Oke",
	}

	userRegister("andre", blacklist)
	userRegister("anjing", blacklist)

	userRegister("DOG", func(name string) bool {
		return name == "DOG"
	})
}

func validation(name string) bool {
	if name == "andre" {
		return true
	} else {
		return false
	}
}
