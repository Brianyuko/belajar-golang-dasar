package main

import "fmt"

func main() {
	brand := "Sony"
	fmt.Println(brand)
	brand = "Microsoft"
	fmt.Println(brand)

	var name string

	name = "Brian Yuko"
	fmt.Println(name)

	name = "Brian Yuko Putra"
	fmt.Println(name)

	var country string = "Indonesia"
	fmt.Println(country)

	var city = "Jakarta"
	fmt.Println(city)

	// Multiple Variable
	var (
		firstName = "Brian"
		lastName  = "Yuko"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
