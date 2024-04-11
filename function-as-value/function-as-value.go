package main

import "fmt"

func getGoodBye(name string) string {
	return "Sayonara " + name
}

func main() {
	sayGoodBye := getGoodBye

	fmt.Println(sayGoodBye("Saske"))

	result := sayGoodBye("Narto")
	fmt.Println(result)
}
