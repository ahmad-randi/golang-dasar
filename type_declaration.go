package main

import "fmt"

func main() {
	type NoKtp string

	var ktpBedul NoKtp = "1234567890"
	fmt.Println(ktpBedul)

	fmt.Println(NoKtp("22222222222222"))
}
