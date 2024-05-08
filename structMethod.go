package main

import (
	"fmt"
)

func main(){
	fmt.Println("Hello Struct")
	rully := Customer{name : "rully", address :"SBY", age : 19}
	rully.sayHello()
}

func (customer Customer)sayHello(){
	fmt.Println("Hello, my name is ", customer.name)
}
type Customer struct{
	name, address string
	age int
}