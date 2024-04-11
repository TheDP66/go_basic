package main

import (
	"fmt"
	"os" // ? Untuk mendapatkan
)

func main() {
	// ? Args
	// ? Mendapatkan Argumen saat menjalankan file ini
	args := os.Args
	fmt.Println("Argument :")
	fmt.Println(args)

	// ! Mendapatkan lokasi binary file
	fmt.Println(args[0])

	// ? Selalu mulai dari index 1
	fmt.Println(args[1])
	fmt.Println(args[2])

	// * Jalankan syntak ini dengan menambahkan argumen
	// go run .\packages\os\os.go Dharma Putra

	// ? Hostname
	hostname, err := os.Hostname()
	fmt.Println("Hostname :")

	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}
