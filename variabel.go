package main

import "fmt"

func main() {

	//deklarasi variabel
	var name, prodi string
	var umur int

	//inisialisasi variabel
	name = "Ahmad Randi"
	umur = 23
	prodi = "Teknik Informatika"
	hobi := "Belajar Coding"

	//menampilkan isi variabel
	fmt.Println(name)
	fmt.Println(umur)
	fmt.Println(prodi)
	fmt.Println(hobi)

	fmt.Println("========================================")
	//deklarasi dan inisialisasi variabel sekaligus
	var (
		namadepan    = "Ahamd"
		namabelakang = "Randi"
		age          = 23
	)

	fmt.Println(namadepan)
	fmt.Println(namabelakang)
	fmt.Println("Nama Saya, ", namadepan, ""+namabelakang)
	fmt.Println(age)
}
