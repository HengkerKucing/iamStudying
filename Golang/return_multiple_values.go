package main

import "fmt"

func getFullName() (string, string) {
	return "luqman", "aldi"
}

func main() {
	// firstName, lastname := getFullName()
	// fmt.Println(firstName, lastname)

	firstName, _ := getFullName()
	fmt.Println(firstName)
}
 