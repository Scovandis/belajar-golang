package main

import "fmt"

func main() {
	fmt.Println("HALLO")
	fmt.Println(factorialLoop(5))
	fmt.Println("RESCURSIVE\n", factorialRescursive(5))
}

func factorialLoop(value int) int {
	total := 1
	for i := value; i > 0; i-- {
		total *= i
	}

	return total
}

func factorialRescursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRescursive(value-1)
	}
}
