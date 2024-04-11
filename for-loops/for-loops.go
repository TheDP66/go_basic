package main

import "fmt"

func main() {
	counter := 1

	for counter <= 5 {
		fmt.Println(counter, "loop")
		counter++
	}

	for idx := 1; idx <= 5; idx++ {
		fmt.Println("Index", idx)
	}

	slice := []string{"Liang", "Yong", "Guang"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for _, item := range slice {
		fmt.Println(item)
	}

	for idx, item := range slice {
		fmt.Println(idx, "-", item)
	}

	person := make(map[string]string)
	person["name"] = "Dharma"
	person["age"] = "21"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
