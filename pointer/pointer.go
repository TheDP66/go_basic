package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// ? data dari address1 akan di clone ke address2
	address2 := address1

	// ? address3 akan pointer ke data address1
	address3 := &address1

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)

	address3.City = "Surabaya"
	fmt.Println(address3)
	fmt.Println(address1)

	// ? Memaksa seluruh pointer address dan address itu sendiri
	// ? Untuk pindah ke address4
	address4 := &address1
	*address4 = Address{"Melbourne", "East", "Australia"}
	fmt.Println(address4)
	fmt.Println(address3)
	fmt.Println(address1)

	// ? Membuat Address yang isinya kososng
	var address5 *Address = new(Address)
	fmt.Println(address5)

	address6 := new(Address)
	address6.City = "Malang"
	fmt.Println(address6)
}
