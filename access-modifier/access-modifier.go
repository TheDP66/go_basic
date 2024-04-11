package main

// ? Tidak bisa export / diakses dari luar karena huruf depan kecil
var version = "0.0.1"

// ? Tidak bisa export / diakses dari luar karena huruf depan kecil
func sayGoodBye(name string) string {
	return "Bye " + name
}

// ? Bisa export / diakses dari luar karena huruf depan besar
var Application = "golang"

// ? Bisa export / diakses dari luar karena huruf depan besar
func SayHello(name string) string {
	return "Hello " + name
}

func main() {

}
