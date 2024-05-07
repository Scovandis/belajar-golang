package main

import (
	"fmt"
)

func main() {
	fmt.Println(recalculate("andre", 100))
	fmt.Println(recalculates("Dre", 999))
	fmt.Println(recalculates("DreSIh", 9))
}

func recalculate(name string, nilai int) string {
	var hasil = ""

	if nilai > 70 && nilai < 80 {
		hasil = "Nilai kamu cukup"
	} else if nilai > 80 && nilai < 90 {
		hasil = "Bagus!!"
	} else if nilai > 90 && nilai <= 100 {
		hasil = "Perfect!! pertahankan."
	} else if nilai < 70 {
		hasil = "yahhh, semangat lagi yahhh!!"
	} else {
		hasil = "Went's wrong to input values"
	}

	return fmt.Sprintf("Hallo %s, %s", name, hasil)
}

func recalculates(name string, nilai int) string {
	var hasil = ""

	switch {
	case nilai < 70:
		hasil = fmt.Sprintf("yahhh nilai kamu kurang, semangat lagi yahh %s!!", name)
	case nilai > 70 && nilai < 80:
		hasil = fmt.Sprintf("nilai kamu cukup %s, semangat", name)
	case nilai > 80 && nilai < 90:
		hasil = fmt.Sprintf("Bagus %s!!", name)
	case nilai > 80 && nilai <= 100:
		hasil = fmt.Sprintf("Perfect, Pertahankan lagi %s", name)
	default:
		hasil = "Went's wrong to input values"
	}

	return fmt.Sprintf("Nilai : %d\n%s", nilai, hasil)
}
