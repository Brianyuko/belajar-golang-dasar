package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var d = 5
	var e = 2
	var c = a + b - d*e // 10
	fmt.Println(c)

	//  Augmented Assignments
	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)
	i += 5
	fmt.Println(i)

	// Unary Operator
	var j = 1
	j++ // j = j + 1
	fmt.Println(j)
	j++
	fmt.Println(j)
}
