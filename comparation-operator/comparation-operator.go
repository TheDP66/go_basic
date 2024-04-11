package main

import "fmt"

func main() {
	name1 := "Eko"
	name2 := "Eko"

	var result bool = name1 == name2
	fmt.Println(result)

	val1 := 100
	val2 := 200
	fmt.Println(val1 > val2)
	fmt.Println(val1 < val2)
	fmt.Println(val1 <= val2)
	fmt.Println(val1 != val2)
}
