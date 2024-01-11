package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari", "Februari", "March", "April", "Mei", "Juni", "Juli", "Agustus", "September", "November", "Desember",
	}

	// var slice1 = months[4:11]

	fmt.Println(months)
	fmt.Println(months[0:2])
	fmt.Println(len(months[0:2]))
	fmt.Println(cap(months))

	var slice1 = months[3:5]
	fmt.Println(fmt.Sprintf("slice 1 %s", slice1))
	slice1[0] = "NIHIL"
	fmt.Println(slice1)
}
