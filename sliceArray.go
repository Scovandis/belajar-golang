package main

import (
	"fmt"
)

// fun convertStringToInt(values : *strconv.T): string{
// 	num, err :=  strconv.Atoi(values)

// 	if err != nil {
// 		fmt.Println("Error ", err)
// 		return "0"
// 	}

// 	fmt.Println("success convert!!")
// 	return num
// }

// func remove[]T(slice []string, index int) int {
// 	return append(slice[:index], slice[index+1]...)
// }

func main() {
	// var months = [...]string{
	// 	"Januari", "Februari", "March", "April", "Mei", "Juni", "Juli", "Agustus", "September", "November", "Desember",
	// }

	// // var slice1 = months[4:11]

	// fmt.Println(months)
	// fmt.Println(months[0:2])
	// fmt.Println(len(months[0:2]))
	// fmt.Println(cap(months))

	// var slice1 = months[3:5]
	// fmt.Println(fmt.Sprintf("slice 1 %s", slice1))
	// slice1[0] = "NIHIL"
	// fmt.Println(slice1)

	// var days = [...]string{
	// 	"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu",
	// }

	// fmt.Println(days)
	// daysSlice := days[3:]
	// // daysSlice[0] = "test aja sih"

	// daysSlice2 := append(daysSlice, "Libur")
	// daysSlice2[0] = "HMMMMM"

	// fmt.Println(len(daysSlice))
	// fmt.Println(fmt.Sprintf("slice 1 : %s", daysSlice))
	// fmt.Println(fmt.Sprintf("slice 2 : %s", daysSlice2))
	// fmt.Println(fmt.Sprintf("Slince length : %s\n cap : %s", len(daysSlice), cap(daysSlice)))

	// fmt.Println(fmt.Sprintf("from : %s change into %s", days[6:], daysSlice))

	// fmt.Println("result convert : ", convertStringToInt("12"))

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	fmt.Println(months)

	slice1 := months[4:6]
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	fmt.Println(slice1)

	// months[5] = "Ubah"
	// fmt.Println(months)

	// slice1[0] = "April Mop"
	// fmt.Println(slice1)

	slice1 = append(slice1, "Ajustment")
	fmt.Println(slice1)

	// slice1 = remove(slice1, 0)
	// fmt.Println(slice1)

}
