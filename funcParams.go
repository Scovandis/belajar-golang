package main

import (
	"fmt"
)

func main() {
	calculate(2, 2)
	calculate(4, 2)
	calculate(1, 5)
}

func calculate(number int, lastNumber int) {
	if number < lastNumber {
		fmt.Println(fmt.Sprintf("Number : %d < %d LastNumber", number, lastNumber))
		// fmt.Println("Number : %s < %s LastNumber", number, lastNumber)
	} else if number > lastNumber {
		fmt.Println(fmt.Sprintf("Number : %d > %d LastNumber", number, lastNumber))
		// fmt.Println("Number : %s > %s LastNumber", number, lastNumber)
	} else if number == lastNumber {
		fmt.Println(fmt.Sprintf("Number : %d == %d LastNumber", number, lastNumber))
		// fmt.Println("Number : %s == %s LastNumber", number, lastNumber)
	} else {
		fmt.Println("wants Wrong")
	}
}
