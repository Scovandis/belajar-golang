package main

import (
	"fmt"
)

func main(){
	fmt.Println("HHH")

	var eko Person
	eko.name = "eko"
	sayHello(eko)

	cat := Animal{name : "Kucing"}
	sayHello(cat)
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

type Animal struct{
	name string
}

func (animal Animal) GetName() string{
	return animal.name
}