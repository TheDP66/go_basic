package main

import "fmt"

func main() {
	var names [5]string
	names[0] = "Liang"
	names[1] = "Young"
	names[2] = "Guang"

	//! Error karena index out of bounds
	// names[3] = "LOL"

	fmt.Println(names[2])

	var number = [5]int{
		5, 15, 30,
	}
	fmt.Println(number)

	// Array Function
	fmt.Println(len(number))

	var word = [7]string{}
	fmt.Println(word)
	fmt.Println(len(word))
}
