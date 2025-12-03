package main

import (
	"fmt"
)

func main() {

	// 1. Membuat array awal
	name := [...]string{"Bedul", "Ahamd", "Randi", "Nabila", "Zahra", "Bella", "Sapina"}

	// Buat slice dari array
	slice := name[1:4] // Ambil Ahamd, Randi, Nabila
	fmt.Println("Slice Awal:", slice)

	// 2. len() dan cap()
	fmt.Println("Len:", len(slice)) // panjang elemen
	fmt.Println("Cap:", cap(slice)) // kapasitas dari index slice pertama ke akhir array

	// 3. append()
	fmt.Println("\n=== Append ===")
	newSlice := append(slice, "Baru Ditambah")
	fmt.Println("Slice Hasil Append:", newSlice)

	// ❗ Catatan penting:
	// Jika append menyebabkan kapasitas penuh → Go buat array baru
	// sehingga perubahan di newSlice tidak mengubah array name

	// 4. make()
	fmt.Println("\n=== make() ===")
	sliceMake := make([]string, 3, 5)
	sliceMake[0] = "Data 1"
	sliceMake[1] = "Data 2"
	sliceMake[2] = "Data 3"
	fmt.Println("Slice dari make:", sliceMake, "Len:", len(sliceMake), "Cap:", cap(sliceMake))

	// 5. copy()
	fmt.Println("\n=== copy() ===")
	src := []string{"A", "B", "C", "D"}
	dst := make([]string, 2) // dst lebih pendek

	copy(dst, src) // hanya 2 data pertama yg masuk
	fmt.Println("Source:", src)
	fmt.Println("Destination:", dst)
}
