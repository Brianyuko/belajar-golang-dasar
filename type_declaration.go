package main

import "fmt"

func main() {
	type IdCard string // Alias for type string

	var idCardBrian IdCard = "111"
	var example string = "22222"
	var exampleIdCard IdCard = IdCard(example)

	fmt.Println(idCardBrian)
	fmt.Println(exampleIdCard)
}
