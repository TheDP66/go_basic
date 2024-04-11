package main

import (
	"fmt"
	"strings"
)

func main() {
	// ? Menghilangkan cutset diawal dan diakhir
	trim := strings.Trim(" abc ", " ")
	fmt.Println(trim)

	lower := strings.ToLower("ABC")
	fmt.Println(lower)

	UPPER := strings.ToUpper("abc")
	fmt.Println(UPPER)

	split := strings.Split("abcbd", "b")
	fmt.Println(split)

	contains := strings.Contains("abc", "b")
	fmt.Println(contains)

	replaceAll := strings.ReplaceAll("abcbdbe", "b", " ")
	fmt.Println(replaceAll)
}
