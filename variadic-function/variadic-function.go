package main

import "fmt"

//? numbers berisikan seluruh sisa parameter yang diberikan ke function
func sumAll(name string, numbers ...int) (string, int) {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return name, total
}

func main() {
	name, total := sumAll("sixsix", 10, 10, 10, 10, 10, 10, 6)
	fmt.Println(name, total)

	numbers := []int{10, 10, 10, 10, 10, 10, 6}
	name2, total2 := sumAll("SIXSIX", numbers...)
	fmt.Println(name2, total2)
}
