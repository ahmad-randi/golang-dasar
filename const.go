package main

import "fmt"

func main() {
	const namaDepan string = "Ahmad"
	const namaBelakang string = "Randi"
	const namaPanggilan string = "Bedul"
	const umur = 23

	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
	fmt.Println(namaPanggilan)
	fmt.Println(umur)

	fmt.Println("==================================")

	const (
		fistName = "Ahmad"
		lastName = "Randi"
		age      = 23
	)

	fmt.Println(fistName)
	fmt.Println(lastName)
	fmt.Println(age)
}
