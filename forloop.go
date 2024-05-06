package main

import (
	"fmt"
)

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke : ", counter)
		counter++
	}

	for number := 1; number <= 10; number++ {
		fmt.Println("number ke : ", number)
	}

	slice := []string{"Nobody", "kornia", "kadom", "fake name", "full name", "kogan"}

	for number := 0; number < len(slice); number++ {
		fmt.Println("Nama : ", slice[number])
	}

	for index := range slice {
		fmt.Println("ini ambil index : ", slice[index])
	}

	for i, names := range slice {
		fmt.Println("index & name, index : ", i, "Names of : ", names)
	}

	for _, names := range slice {
		fmt.Println("_, Names of : ", names)
	}

	person := map[string]string{
		"name":    "nobody",
		"address": "fall",
	}

	for index, name := range person {
		fmt.Println("Index :", index, "person : ", name)
	}
}
