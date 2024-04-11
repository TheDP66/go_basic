package main

import "fmt"

// ? interface biasa digunakkan untuk membuat contract

// 1. Deklarasi interface
type HasName interface {
	GetName() string
}

// 2. Buat function dengan parameter interface
func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// 3. Buat Struct
type Person struct {
	Name string
}

// 4. Buat Struct dengan return yang sesuai dengan interface
// ? Nama function, parameter, dan return harus sama dengan interface
// ? Jika tidak maka akan error
func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	// 5. Deklarasi variable ke Struct
	dharma := Person{
		Name: "Dharma",
	}

	// 6. Panggil function
	SayHello(dharma)

	cat := Animal{
		Name: "Puss",
	}

	SayHello(cat)
}
