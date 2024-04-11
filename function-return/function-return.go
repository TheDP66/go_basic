package main

import "fmt"

func getGreet(name string) string {
	return "Haiii " + name
}

func main() {
	result := getGreet("Dharma")

	fmt.Println(result)
}
