package main

import "fmt"

func main() {
	var name [3]string
	name[0] = "Bedul"
	name[1] = "Ahmad"
	name[2] = "Randi"
	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [...]int{
		90,
		80,
		70,
	}
	fmt.Println(values)
	//len fungsi untuk menghitung panjang array
	fmt.Println(len(values))
	values[2] = 100
	fmt.Println(values)
}
