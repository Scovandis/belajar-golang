package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 10; i++ {
		// fmt.Println("Perulangan ke : ", i)
		calculate(i)
	}
}

func hello() {
	fmt.Println("Hello World!")
}

func calculate(number int) {
	fmt.Println("Perulangan ke : ", number)
}
