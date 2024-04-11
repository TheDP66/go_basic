package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	//? Cara 1: Function anonymous dibuat terpisah
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUser("dharma", blacklist)

	//? Cara 2: Function anonymous ditulis di parameter
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}

//? Function tanpa memiliki nama
