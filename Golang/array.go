package main

import "fmt"


func main() {
	var name [3]string

	name[0] = "luqman"
	name[1] = "aldi"
	name[2] = "prawiratama"

	fmt.Println(name[0])
	fmt.Println(name[1])

	var saya = [...]int {
		90, 80, 40 }

	fmt.Println(saya)
	fmt.Println(len(saya))
}