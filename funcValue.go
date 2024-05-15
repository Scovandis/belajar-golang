package main

import (
	"fmt"
)

func main() {

	goodBye := getGoodBye

	fmt.Println(goodBye("andre"))
}

func getG‹oodBye(name string) string {
	return fmt.Sprintf("Selamat tinggal %s!!, sampai bertemu di lain hari.", name)
}
