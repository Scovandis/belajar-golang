package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	result := random()
	resultString := result.(string)
	fmt.Println("result :", resultString)

	/* 	resultInt := result.(int)
	   	fmt.Println("resultInt :", resultInt) */

	getAssert()
}

func random() interface{} {
	return "Oke"
}

func getAssert() {
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("String ", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println()
	}
}
