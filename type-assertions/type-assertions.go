package main

import "fmt"

func random() interface{} {
	return 30
}

func main() {
	result := random()
	// ! Error karena tipe tidak sesuai
	// resultString := result.(int)

	// Cara 1 : Tidak aman (Bisa Error)
	// resultString := result.(string)
	// fmt.Println(resultString)

	// Cara 2 : Aman (Menggunakan Switch)
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown")
	}
}
