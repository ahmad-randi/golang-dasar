package main

import "fmt"

func main() {
	name := "dul"

	switch name {
	case "Bedul":
		fmt.Println("Hai", name)
	case "Ahmad":
		fmt.Println("Hai", name)
	case "Randi":
		fmt.Println("Hai", name)
	default:
		fmt.Println("Hai", name, "Boleh kita kenalan?")
	}

	//Switch dengan Short Statement
	switch length := len(name); length > 3 {
	case false:
		fmt.Println("Nama terlalu panjang. Panjang nama anda", length)
	case true:
		fmt.Println("Nama anda memenuhi sarat panjang karakter", length)
	}
}
