package main

import "fmt"

func getHello(nama string) string {
	hello := "hello " + nama

	return hello
}

func main () {
	var result = getHello("aldi")
	fmt.Println(result)
}