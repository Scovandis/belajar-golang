package main

import (
	"fmt"
)

func main(){
	fmt.Println("Hello")

	var customer Customer
	customer.name = "Andre"
	customer.address = "Surabaya"
	customer.age = 99

	fmt.Println(customer)

	datas := Customer{
		name : "Wahyono",
		address : "Sidoearjo",
		age : 11,
	}

	fmt.Println(datas)

	// not recommended
	budi := Customer{"Budi", "semarang", 12,}
	fmt.Println(budi)
}

type Customer struct {
	name, address string
	age int
}