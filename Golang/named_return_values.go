package main

import "fmt"

func getCompleteName() (firstName string, middleName string, lastname string) {
	firstName = "luqman"
	middleName = "aldi"
	lastname = "prawiratama"

	return firstName, middleName, lastname
}

func main() {
	a, b, c := getCompleteName()

	fmt.Println(a)
}