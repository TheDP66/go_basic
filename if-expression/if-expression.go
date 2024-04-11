package main

import "fmt"

func main() {
	status := 402

	if status == 200 {
		fmt.Println("Success")
	} else if status == 404 {
		fmt.Println("Not Found")
	} else {
		fmt.Println("Unknown")
	}

	var otp = "1234"

	//! Cara Panjang
	// var length = len(otp)
	// if length > 4 {

	//? Cara Singkat
	if length := len(otp); length > 4 {
		fmt.Println("Terlalu Panjang")
	} else if length == 4 {
		fmt.Println("Sudah Benar")
	} else {
		fmt.Println("Terlalu Pendek")
	}
}
