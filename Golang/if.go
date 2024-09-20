package main

import "fmt"

func main() {
	name := "aldi"

	if name == "aldi" {
		fmt.Println("halo aldi")
	} else if name == "jokondo kondo" {
		fmt.Println("hallo jokondo kondo")
	} else {
		fmt.Println("hai boleh kenalan ngga kiw kiw")
	}

	
	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")

	} else {
		fmt.Println("nama sudah ok")
	}
}
