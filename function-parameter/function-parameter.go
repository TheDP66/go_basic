package main

import "fmt"

func sayHelloTo(firstName string, lastname string) {
	fmt.Println("Heij,", firstName, lastname)
}

func main() {
	sayHelloTo("Dharma", "Putra")
}
