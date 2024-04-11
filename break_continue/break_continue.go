package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 6 {
			fmt.Println("break")
			break
		}

		fmt.Println(i, "loops")
	}

	for j := 0; j < 10; j++ {
		if j == 6 {
			fmt.Println("continue")
			continue
		}

		fmt.Println(j, "loops")
	}
}
