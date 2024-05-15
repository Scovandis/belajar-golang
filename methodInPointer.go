package main

import (
	"fmt"
)

type Men struct {
	name string
}

func (men *Men) Merried() {
	men.name = "Mr." + men.name
}
func main() {
	andre := Men{"Andre"}
	andre.Merried()
	fmt.Println("Helo", andre.name)
}
