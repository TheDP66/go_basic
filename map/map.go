package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Dharma",
		"city": "Surabaya",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["city"])

	person["title"] = "Programmer"
	fmt.Println(person["title"])

	fmt.Println(len(person))

	//? Menghapus suatu key di map
	delete(person, "city")
	fmt.Println(person)

	fmt.Println("================================================================")

	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Not Me"
	book["ups"] = "Salah"
	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)
}
