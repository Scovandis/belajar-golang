package main

import (
	"fmt"
)

func main() {
	fmt.Println("HEllo")

	ups := Ups(5)
	var upss interface{} = Ups(3)
	fmt.Println(ups)
	fmt.Println(upss)
}

func Ups(i int) interface{} {
	if i == 1 {
		return true
	} else if i == 2 {
		return false
	} else if i == 3 {
		return "Ups"
	} else {
		return i
	}
}
