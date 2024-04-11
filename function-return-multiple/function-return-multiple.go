package main

import (
	"fmt"
	"strings"
)

func splitName(fullName string) (string, string) {
	if strings.Contains(fullName, " ") {
		return strings.Split(fullName, " ")[0], strings.Split(fullName, " ")[1]
	} else {
		return fullName, ""
	}

}

func main() {
	firstName, lastName := splitName("Dharma Putra")

	fmt.Println("firstName", firstName)
	fmt.Println("lastName", lastName)

	firstName2, lastName2 := splitName("Guang")

	fmt.Println("firstName2", firstName2)
	fmt.Println("lastName2", lastName2)
}
