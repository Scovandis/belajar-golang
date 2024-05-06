package main

import (
	"fmt"
)

func main() {
	var name = "andre"
	age := 12

	if name != "andre" && age == 12 {
		fmt.Println("Hallo NPC")
	} else if name == "andre" && age != 12 {
		fmt.Println("OKE")
	} else {
		fmt.Println("Flase")
	}

	if len(name) > 4 {
		fmt.Println("pass")
	} else {
		fmt.Println("wrong")
	}
}
