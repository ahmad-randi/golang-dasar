package main

import "fmt"

func main() {
	fmt.Println("Nama = Ahmad Randi")
	fmt.Println("Nama Depan = Ahmad ")
	fmt.Println("Nama Belakang = Randi ")
	fmt.Println("Nama Panggilan = Bedul ")

	fmt.Println(len("Bedul"))

	//conversi string
	var name = "Bedul"
	var B = name[0]
	var E = name[1]
	var D = name[2]
	var U = name[3]
	var L = name[4]
	var b = string(B)
	var e = string(E)
	var d = string(D)
	var u = string(U)
	var l = string(L)
	fmt.Print(b, e, d, u, l)
	fmt.Println()

	//conversi cepat
	fmt.Println("Convert cepat dengan for")
	var name2 = "Bedul"
	for _, c := range name2 {
		fmt.Print(string(c))
	}
}
