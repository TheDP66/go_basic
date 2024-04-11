package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hi", name, ", My name is", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuu", a.Name)
}

func main() {
	liang := Customer{
		Name:    "Liang Yong Guang",
		Address: "Surabaya",
		Age:     21,
	}

	liang.sayHi("Dharma")
	liang.sayHuuu()
}
