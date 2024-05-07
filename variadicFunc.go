package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello\n\n")

	total := summAll(10, 20, 30, 40)
	fmt.Println("total ", total)

	fmt.Println(recalculate("Dre", 10, 20, 30, 40))

	numbers := []int{100, 10, 1}
	totals := summAll(numbers...)
	fmt.Println(totals)

}

func summAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func recalculate(name string, numbers ...int) string {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return fmt.Sprintf("Hallo %s!, summAll adalah %d", name, total)
}
