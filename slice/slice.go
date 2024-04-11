package main

import "fmt"

func main() {
	//? Kalau tidak tau length arraynya bisa diisi dengan ...
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	fmt.Println(months)
	fmt.Println(len(months))
	fmt.Println("================================================")

	//? Slice dari array asal
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(slice1[0])

	//? len() Mendapat panjang slice
	fmt.Println(len(slice1))

	//? cap() Mendapat kapasitas sesungguhnya slice
	fmt.Println(cap(slice1))

	fmt.Println("================================================")

	//? Saat slice diubah array asal juga ikut terubah
	//? Saat array diubah slice juga ikut terubah
	slice1[0] = "January"
	fmt.Println(months)

	fmt.Println("================================================")

	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days)

	fmt.Println("================================================")

	//? append() Menambahkan data baru dengan membuat slice baru
	daysSlice2 := append(daySlice1, "Libur YEY")
	daysSlice2[0] = "Ups"
	fmt.Println(daysSlice2)
	fmt.Println(days)

	fmt.Println("================================================")

	//? make() Membuat slice baru dari awal
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Dharma"
	newSlice[1] = "Putra"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fmt.Println("================================================")

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	fmt.Println("================================================")

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
