package main

import "fmt" 

func main() {

	type NoKTP string

	var Aldi NoKTP = "123456"
	var contoh string =  "0000"

	var contohKTP NoKTP = NoKTP(contoh)


	fmt.Println(Aldi)
	fmt.Println(contoh)
	fmt.Println(contohKTP)
}