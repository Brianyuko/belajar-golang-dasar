package main

import (
	"fmt"
)

func main() {
	var x = 5
	var y = 3

	var name1 = "Brian"
	var name2 = "Brian"

	var result bool = name1 == name2
	fmt.Println(x > y)  // returns 1 (true) because 5 is greater than 3
	fmt.Println(result) // returns 1 (true) because name1 is equal to name2
}
