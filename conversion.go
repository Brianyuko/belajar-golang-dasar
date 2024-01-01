package main

import "fmt"

func main() {
	var value32 int32 = 32768
	var value64 int64 = int64(value32)
	var value16 int16 = int16(value32) // Max Value 32767 & Min Value -32768

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value16) //Overflow

	name := "Brian Yuko"
	e := name[0] //Type  uint8
	eString := string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
