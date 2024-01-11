package main

import "fmt"

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

	var days = [...]string{
		"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu",
	}

	fmt.Println(days)
	daysSlice := days[3:]
	daysSlice[0] = "test aja sih"

	daysSlice2 := append(daysSlice, "Libur")
	daysSlice2[0] = "HMMMMM"

	fmt.Println(fmt.Sprintf("slice 1 : %s", daysSlice))
	fmt.Println(fmt.Sprintf("slice 2 : %s", daysSlice2))

	fmt.Println(fmt.Sprintf("from : %s change into %s", days[6:], daysSlice))

}
