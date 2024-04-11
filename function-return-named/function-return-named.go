package main

import "fmt"

func getFullName() (firstName string, middleName string, lastName string) {
	firstName = "Liang"
	middleName = "Yong"
	lastName = "Guang"
	return
}

func main() {
	f, m, l := getFullName()
	fmt.Println(f, m, l)
}
