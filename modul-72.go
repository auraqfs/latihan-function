package main

import (
	"fmt"
	"sort"
	"strings"
)

type listMusik struct {
	Judul, penyanyi string
	jmlVote         float64
}

var jdl, nama string
var vote float64

var id int = 0

var list = []listMusik{}

func addVote() {
	fmt.Print("\nMasukkan Judul Lagu : ")
	fmt.Scan(&jdl)
	fmt.Print("Masukkan Nama Penyanyi : ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan Jumlah Vote (1.0-5.0) : ")
	fmt.Scan(&vote)

	list = append(list, listMusik{
		Judul:    jdl,
		penyanyi: nama,
		jmlVote:  vote,
	})

	sort.Slice(list[:], func(i, j int) bool {
		return list[i].jmlVote > list[j].jmlVote
	})
}

func deleteData() {

	if len(list) == 0 {
		fmt.Println("Tidak ada data yang bisa dihapus")
	}
	var before = list[:id-1]
	var after = list[id:]
	list = append(before, after...)

}

func viewTop3() {
	var top int = 0
	var i float64 = 5.0

	for i >= 0 {
		for j := range list {
			lists := list[j]
			if lists.jmlVote <= i {
				fmt.Println(top+1, "\t", lists.Judul, "\t\t\t", lists.penyanyi, "\t\t\t", lists.jmlVote)
				top++
			}
			if top == 3 {
				break
			}
		}
		if top <= 3 {
			break
		}
	}
}

func searchData() {
	for j := range list {
		lists := list[j]
		var isPrefix = strings.HasPrefix(lists.penyanyi, "a")
		if isPrefix == true {
			fmt.Println(j+1, "\t", lists.Judul, "\t\t\t", lists.penyanyi, "\t\t\t", lists.jmlVote)
		} else {
			fmt.Println("\n\t\t\tDATA TIDAK DITEMUKAN!")
		}
	}
}

func main() {
	var pilih int
	var x = true

	for x {
		fmt.Println("\n\t\t\t====Daftar Pilihan==== ")
		fmt.Println("1. Input Data Voting")
		fmt.Println("2. Hapus data vote berdasarkan id ")
		fmt.Println("3. Tampilkan seluruh data vote beserta jumlah data yang tersimpan dalam list")
		fmt.Println("4. Menampilkan jumlah keseluruhan jumlah vote ")
		fmt.Println("5. Menampilkan top 3 musik terfavorit")
		fmt.Println("6. Menampilkan seluruh data dengan nama penyanyi berawalan huruf A")
		fmt.Println("0. Keluar")
		fmt.Println("=================================================================== ")
		fmt.Print("Masukkan Pilihan : ")
		fmt.Scan(&pilih)
		fmt.Println("==================================================================== ")

		switch pilih {
		case 1:
			addVote()

		case 2:
			fmt.Print("\nMasukkan ID yang akan dihapus: ")
			fmt.Scan(&id)

			deleteData()
			sort.Slice(list[:], func(i, j int) bool {
				return list[i].jmlVote > list[j].jmlVote
			})
		case 3:

			fmt.Println("ID \tJudul Lagu \t\tPenyanyi \t\tJumlah Vote")

			for i := range list {
				lists := list[i]
				fmt.Println(i+1, "\t", lists.Judul, "\t\t\t", lists.penyanyi, "\t\t\t", lists.jmlVote)
			}
			fmt.Println("==================================================================== ")
			fmt.Println("Jumlah Data Dalam List : ", len(list), "Item")

		case 4:
			var jumlah float64 = 0.0
			for i := range list {
				lists := list[i]
				jumlah += lists.jmlVote
			}

			fmt.Print("Jumlah Keseluruhan Jumlah Vote : ")
			fmt.Println(jumlah)

		case 5:
			fmt.Println("                        TOP 3 MUSIK TERFAVORIT                       ")
			fmt.Println("==================================================================== ")
			fmt.Println("ID \tJudul Lagu \t\tPenyanyi \t\tJumlah Vote")
			viewTop3()

		case 6:
			fmt.Println("            DATA DENGAN NAMA PENYANYI BERAWALAN HURUF A              ")
			fmt.Println("==================================================================== ")
			fmt.Println("ID \tJudul Lagu \t\tPenyanyi \t\tJumlah Vote")
			searchData()

		case 0:
			x = false
		}
	}
}
