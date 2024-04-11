package main

import "fmt"

func main() {
	var name string

	name = "Dharma Putra"
	fmt.Println(name)

	name = "Liang Young"
	fmt.Println(name)

	//! Error karena tipe data beda
	// name = true

	//? Jika nilai variabel didefine dari awal
	//? user tidak perlu menginputkan tipe data
	var age = 14
	fmt.Println(age)

	//? Atau bisa diberi tipe jika mau dipaksa
	var height int16 = 170
	fmt.Println(height)

	//? kata kunci var bisa dihilangkan dengan :=
	country := "Indonesia"
	fmt.Println(country)

	country = "Japan"

	//? Jika mau declare lebih dari satu variabel
	var (
		firstName = "Liang"
		lastName  = "Young"
	)

	fmt.Println(firstName, lastName)

}
