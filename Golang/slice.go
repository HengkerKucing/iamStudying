package main

import "fmt"

func main() {
	nama := [...]string {"Luqman", "Aldi", "Prawiratama", "Salsa", "Anica", "Septaningtias"}

	slice1 := nama[1:2]
	fmt.Println(slice1)
	slice2 := nama[:3]
	fmt.Println(slice2)
	slice3 := nama[4:]
	fmt.Println(slice3)
	slice4 := nama[:]
	fmt.Println(slice4)


	days := [...]string {"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	fmt.Println(days)
	daySlice1 := days[5:]
	fmt.Println(daySlice1)
	daySlice1[0] = "halodek"
	daySlice1[1] = "halodekkk"
	fmt.Println(daySlice1)
	fmt.Println(days)


	daySlice2 := append(daySlice1, "libur baru")
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)


	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "aldi"
	newSlice[1] = "luqman"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "prawiratama")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))
	fmt.Println(newSlice)

	newSlice2[1] = "ica"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)


}