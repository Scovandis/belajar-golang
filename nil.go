package main

import (
	"fmt"
)

func main() {
	fmt.Println("Helllo")
	name := NewMap("Andre")

	fmt.Println(name)

	var person map[string]string = nil

	if person == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(person)
	}

}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
