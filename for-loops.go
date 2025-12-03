package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Index ke-", counter)
		counter++
	}

	fmt.Println("===============================")

	//For dengan Statement
	for counter := 0; counter <= 10; counter++ {
		fmt.Println("Index ke-", counter)
	}

	fmt.Println("===============================")

	// For Range
	// ● For bisa digunakan untuk melakukan iterasi terhadap semua data collection
	// ● Data collection contohnya Array, Slice dan Map

	countName := []string{
		"Bedul", "Ahmad", "Randi", "Nabila", "Zahra",
	}
	for index, name := range countName {
		fmt.Println("Index", index, "=", name)
	}
}
