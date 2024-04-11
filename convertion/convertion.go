package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println("nilai32: ", nilai32)
	fmt.Println("nilai64: ", nilai64)
	fmt.Println("nilai16: ", nilai16)
	fmt.Println("nilai8: ", nilai8)

	var name = "Dharma Putra"
	var n byte = name[0]
	var nString = string(n)

	fmt.Println(name)
	fmt.Println(n)
	fmt.Println(nString)
}
