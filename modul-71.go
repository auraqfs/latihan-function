package main

import (
	"fmt"
	"math"
)

var total = 0
var sum = 0
var jmlh = 0.0
var jumlahI = 0.0
var hasil int

func deretAngka(jumlah int, awal int) int {

	if jumlah <= 1 {
		fmt.Printf("%d + ", awal)
		total = total + awal
		return awal
	}

	i := deretAngka(jumlah-1, awal)
	pangkat := int(math.Pow(float64(10), float64(jumlah-1)))
	total = awal * pangkat
	fmt.Printf("%d + ", total)

	return i + total
}

func deretPangkat(kali bool, jenis string, jumlah int, iterasi int, pola int) int {
	pangkat := int(math.Pow(float64(pola), float64(iterasi)))

	if kali {
		if jenis == "pangkat" {
			hasil = pangkat * (iterasi + 1)
		}
	} else {
		hasil = pangkat
	}

	if iterasi < jumlah {
		fmt.Print(hasil, " + ")

		return hasil + deretPangkat(kali, jenis, jumlah, iterasi+1, pola)
	} else {
		fmt.Print(hasil)
		return hasil
	}
}

func main() {
	var jumlah, pilihan int
	var x = true

	for x {
		fmt.Println("\n================== PILIHAN ===================")
		fmt.Println("1. Deret Pangkat 6")
		fmt.Println("2. Deret Perkalian Angka 5")
		fmt.Println("3. Deret Pangkat 7")
		fmt.Println("4. Deret Pangkat ")
		fmt.Println("5. Deret Perkalian Angka 5")
		fmt.Println("0. Keluar")
		fmt.Println("==============================================")
		fmt.Print("Masukkan Pilihan : ")
		fmt.Scan(&pilihan)
		fmt.Println("==============================================")

		switch pilihan {
		case 1:
			fmt.Print("Masukan jumlah deret : ")
			fmt.Scan(&jumlah)
			fmt.Print("SUM = ")
			fmt.Println(" =", deretPangkat(false, "pangkat", jumlah, 0, 6))
		case 2:
			fmt.Print("Masukan jumlah deret : ")
			fmt.Scan(&jumlah)
			fmt.Print("SUM = ")
			fmt.Println(" =", deretAngka(jumlah, 5))
		case 3:
			fmt.Print("Masukan jumlah deret : ")
			fmt.Scan(&jumlah)
			fmt.Print("SUM = ")
			fmt.Println(" =", deretPangkat(false, "pangkat", jumlah, 0, 7))
		case 4:
			fmt.Print("Masukan jumlah deret : ")
			fmt.Scan(&jumlah)
			fmt.Print("SUM = ")
			fmt.Println(" =", deretPangkat(true, "pangkat", jumlah, 0, 10))
		case 5:
			fmt.Print("Masukan jumlah deret : ")
			fmt.Scan(&jumlah)
			fmt.Print("SUM = ")
			fmt.Println(" =", deretPangkat(false, "pangkat", jumlah, 0, 6))

		case 0:
			x = false
		}
	}
}
