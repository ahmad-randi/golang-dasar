package main

import "fmt"

func getFullName() (string, string, string) {
	return "Ahmad", "Randi", "bedul"
}

func main() {
	namaDepan, namaBelakang, namaPanggilan := getFullName()
	fmt.Println(namaDepan, namaBelakang)
	fmt.Println(namaPanggilan)
}
