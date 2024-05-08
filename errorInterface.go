package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Helow")
	var contohError error = errors.New("Ups error")
	fmt.Println(contohError.Error())

	hasil, err := Pembagian(100, 0)

	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println("Error :", err.Error())
	}
}

type error interface {
	Error() string
}

func Pembagian(nilai int, pembagian int) (int, error) {
	if pembagian == 0 {
		return 0, errors.New("Pembagian dengan nol (0)")
	} else {
		return nilai / pembagian, nil
	}
}
