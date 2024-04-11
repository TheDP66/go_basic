package main

import "fmt"

func endApp() {
	// Menangkap pesan error dari panic
	// ? recover disimpan didalam function defer
	message := recover()

	// Jika terdapat error
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}

	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	// Menjalankan fungsi meskipun program error
	// Harus ditaruh diatas / dipanggil awal
	defer endApp()

	if error {
		// Menghentikan program untuk logging error
		panic("APLIKASI ERROR")
	}

	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(true)

	fmt.Println("Setelah runApp")
}
