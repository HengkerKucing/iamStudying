package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}

	// person["addres"] = "Pati",
	// person["name"] = "Aldi",

	person := map[string]string{
		"nama":    "aldi",
		"address": "pati",
	}

	fmt.Println(person["nama"])
	fmt.Println(person["address"])
	fmt.Println(len(person))


	book := make(map[string]string)
	book["tittle"] = "buku golang"
	book["author"] = "aldi"
	book["ups"] = "salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}