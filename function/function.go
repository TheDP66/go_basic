package main

import "fmt"

func main() {
	sayHello()

	for i := 0; i < 3; i++ {
		getId(i)
	}
}

func sayHello() {
	fmt.Println("Hello, world!")
}

func getId(id int) {
	fmt.Println("Id", id)
}
