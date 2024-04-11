package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	// ? Membuat Struct
	// Cara 1
	var liang Customer
	liang.Name = "Liang Yong Guang"
	liang.Address = "Surabaya"
	liang.Age = 21

	fmt.Println(liang)

	// Cara 2
	yong := Customer{
		Name:    "Liang Yong Guang",
		Address: "Surabaya",
		Age:     21,
	}

	fmt.Println(yong)

	// Cara 3
	// ! Tidak rekomen
	// Harus di urutkan berdasarkan implementasinya
	guang := Customer{"Liang Yong Guang", "Surabaya", 21}

	fmt.Println(guang)
}
