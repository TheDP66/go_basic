package main

import "fmt"

func main() {
	name := "Yong"

	//? Switch normal
	switch name {
	case "Liang":
		fmt.Println("Hello")
	case "Yong":
		fmt.Println("Hi")
	default:
		fmt.Println("Konichiwa")
	}

	//? Switch dengan kondisi
	switch length := len(name); length > 3 {
	case true:
		fmt.Println("> 3")
	case false:
		fmt.Println("<= 3")
	}

	length2 := len(name)

	//? Switch sebagai if
	switch {
	case length2 > 3:
		fmt.Println("> 3")
	case length2 <= 3:
		fmt.Println("<= 3")
	default:
		fmt.Println("NICE")
	}
}
