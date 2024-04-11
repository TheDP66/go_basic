package main

import (
	"fmt"
	"strconv"
)

// ? Untuk melakukan conversi dari string ke lainnya ataupun sebaliknya

func main() {
	// ? Setiap fungsi selalu memiliki return value dan errornya
	boolean, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	fmt.Println(strconv.ParseFloat("1.4", 64))
	fmt.Println(strconv.ParseInt("1000000", 10, 64))

	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatInt(4, 2))

	// ? Konversi string ke int langsung
	fmt.Println(strconv.Atoi("20000"))

	// ? Konversi int ke string langsung
	fmt.Println(strconv.Itoa(20000))
}
