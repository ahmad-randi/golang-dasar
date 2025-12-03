package main

import "fmt"

func main() {
	person := map[string]string{
		"Name":   "Bedul",
		"Age":    "23",
		"Alamat": "Mampang Prapatan",
	}

	fmt.Println(person["Name"])
	fmt.Println(person["Age"])
	fmt.Println(person["Alamat"])

	fmt.Println("============================")

	delete(person, "Alamat")
	fmt.Println(person)

	book := make(map[string]string)
	book["Title"] = "Buku Golang"
	book["Author"] = "Ahmad Bedul"
	book["Salah"] = "Tidak ada"

	delete(book, "Salah")
	fmt.Println(book)
}
