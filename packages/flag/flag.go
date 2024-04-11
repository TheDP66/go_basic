package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "user", "Put your database username")
	password := flag.String("password", "pass", "Put your database password")

	// * Wajib dipanggil saat menggunakan flag
	flag.Parse()

	fmt.Println("host: ", *host)
	fmt.Println("username: ", *username)
	fmt.Println("password: ", *password)

	// ? Dijalankan dengan cara
	// ? go run .\packages\flag\flag.go -host=hostnya -username=usernamenya -password=passwordnya

	// ? Jika salah satu flag tidak ditemukan maka akan menggunakan nilai default
}
