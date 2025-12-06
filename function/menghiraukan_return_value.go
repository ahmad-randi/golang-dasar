package main

import "fmt"

func getFullName1() (string, string, string) {
	return "Ahmad", "Randi", "Bedul"
}

func main() {
	_, _, namaPanggilan := getFullName1()
	fmt.Println(namaPanggilan)
}
