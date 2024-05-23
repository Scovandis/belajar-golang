package main

import (
	"fmt"
)

type Address struct {
	provincie, city, country string
}

func main() {
	address1 := Address{
		provincie: "Jawa timur",
		city:      "Surabaya",
		country:   "Indonesia",
	}

	address2 := address1
	address2.city = "Sidoarjo"

	var address3 *Address = &address2
	address3.city = "Mojokerjo"

	var address4 *Address = &address1
	address4 = &Address{
		provincie: "Jawa barat",
		city:      "Bandung",
		country:   "Indonesia",
	}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)
	fmt.Println("HellloWord")

	address6 := &address1

	var address5 *Address = &address1
	*address5 = Address{
		provincie: "Jawa tengah",
		city:      "smr",
		country:   "Indonesia",
	}

	fmt.Println(address1)
	fmt.Println(address5)
	fmt.Println(address6)

	//new
	var address7 *Address = new(Address)
	address7.city = "Koo"
	address7.provincie = "IIIO"
	address7.country = "KJIJID"
	fmt.Println(address7)
}
