package main

import (
	"fmt"
)

func main() {
	var name = "andre"

	switch name {
	case "and":
		fmt.Println("And")
	case "andre":
		fmt.Println("andre")
	default:
		fmt.Println("Wrong")
	}

	switch length := len(name); length > 3 {
	case true:
		fmt.Println("Nama benar")
	case false:
		fmt.Println("Nama salah")
	}

	leng := len(name)
	switch {
	case leng > 3:
		fmt.Println("Nama oke")
	case leng <= 3:
		fmt.Println("Nama salah")
	}
}
