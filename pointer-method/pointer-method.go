package main

import "fmt"

type Man struct {
	Name string
}

// ! Gagal karena parameter tidak pointer
// func (man Man) Married() {
// ? Tambahkan pointer ke parameter
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	dharma := Man{"Dharma"}
	fmt.Println(dharma)

	dharma.Married()
	fmt.Println(dharma)
}
