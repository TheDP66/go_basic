package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// ! Data tidak terubah karena yang dikirim data duplikatnya
// func ChangeAddressToIndonesia(address Address) {
// ? Data akan berubah karena di parameter diberikan tanda (*) pointer
func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{"Surabaya", "Jawa Timur", "Australia"}
	fmt.Println(address)

	// ChangeAddressToIndonesia(address)
	// ? Jika parameternya pointer maka tambahkan (&) di depannya
	ChangeAddressToIndonesia(&address)
	fmt.Println(address)

	// ? Atau bisa seperti ini
	address2 := &Address{"Jakarta", "Jawa Tengah", "Australia"}
	ChangeAddressToIndonesia(address2)
	fmt.Println(address2)

	address3 := new(Address)
	fmt.Println(address3)
}
