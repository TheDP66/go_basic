package main

import "fmt"

// ? Dapat menerima tipe data apapaun dengan interface{}
func Ups(i interface{}) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	kosong := Ups(1)
	fmt.Println(kosong)

	var data interface{} = Ups("LOL")
	fmt.Println(data)
}
