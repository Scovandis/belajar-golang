package main

import (
	"fmt"
)

type Address struct {
	name, city, address string
}

func ChangeAddressToIndonesia(address *Address) {
	address.address = "Changedong"
}
func main() {
	address := Address{"Andre", "Surabaya", ""}
	ChangeAddressToIndonesia(&address)
	fmt.Println(address)
}
