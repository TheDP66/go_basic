package main

import "fmt"

// Cara 1 : Cara biasa menggunakan for loop
func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

// Cara 2 : Cara function recursive
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	// result := factorialLoop(5)
	// fmt.Println(result)

	result2 := factorialRecursive(5)
	fmt.Println(result2)
}
