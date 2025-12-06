package main

import "fmt"

func getHello(nama string) string {
	return "Hello " + nama
}

func main() {
	result := getHello("Bedul")
	fmt.Println(result)
}
