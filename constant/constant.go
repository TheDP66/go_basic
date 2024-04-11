package main

import "fmt"

func main() {
	const firstname string = "Dharma"
	const lastname = "Putra"

	fmt.Println(firstname + " " + lastname)

	//! Akan error jika diubah
	// firstname = "Eko"

	const (
		name = "LOL"
		age  = 15
	)
}
