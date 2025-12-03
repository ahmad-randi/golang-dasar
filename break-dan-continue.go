package main

import "fmt"

// ● Break & continue adalah kata kunci yang bisa digunakan dalam perulangan
// ● Break digunakan untuk menghentikan seluruh perulangan
// ● Continue adalah digunakan untuk menghentikan perulangan yang berjalan, dan langsung
// melanjutkan ke perulangan selanjutnya

func main() {
	//Kode Program Break
	count := 10
	for i := 0; i < count; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Index ke-", i)
	}

	fmt.Println("===============================")

	//Kode Program Continue
	for i := 0; i < count; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Index ke-", i)
	}
}
