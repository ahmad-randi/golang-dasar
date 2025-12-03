package main

import "fmt"

func main() {
	name := "dul"

	//If Expression
	if name == "Randi" {
		fmt.Println("Hai", name)
	}

	fmt.Println("=================================")

	//Else Expression
	if name == "Randi" {
		fmt.Println("Hai", name)
	} else {
		fmt.Println("Boleh kenalan", name, "?")
	}

	fmt.Println("=================================")

	//Else If Expression
	if name == "Randi" {
		fmt.Println("Hai", name)
	} else if name == "Ahmad" {
		fmt.Print("Hai", name)
	} else {
		fmt.Println("Anda bukan Ahmad Randi...")
		fmt.Println("Boleh kita kenalan?, ", name)
	}

	fmt.Println("=================================")

	//If dengan Short Statement
	if length := len(name); length > 3 {
		fmt.Println("Nama terlalu panjang. Panjang nama anda", length)
	} else {
		fmt.Println("Nama anda memenuhi sarat panjang karakter", length)
	}
}
