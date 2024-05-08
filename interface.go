package main

import (
	"fmt"
)

func main(){
	fmt.Println("HHH")

	var eko Person
	eko.name = "eko"
	sayHello(eko)
}

func sayHello(hasName HasName){
	fmt.Println("Hallo, ", hasName.GetName())
}

type HasName interface{
	GetName() string
}

type Person struct{
	name string
}

func (person Person) GetName() string{
	return person.name
}