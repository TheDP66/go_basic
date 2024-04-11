package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpEko NoKTP = "1234567890"
	fmt.Println(noKtpEko)

	var marriedStatus Married = true
	fmt.Println(marriedStatus)
}
