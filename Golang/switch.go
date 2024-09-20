package main

import "fmt"

func main() {
	nama := "aldi"

	switch nama {
	case "aldi":
		fmt.Println("halo aldibek")
	case "eko":
		fmt.Println("halo eko")
	case "orang":
		fmt.Println("halo orang")
	default:
		fmt.Println("boleh kenalan ngga dek")
	}
	
	
	switch length := len(nama); length > 5 {
	case true :
		fmt.Println("nama kepanjangan")
	case false :
		fmt.Println("nama sudah ok")
	}

	length := len(nama)
	switch {
	case length > 5 :
		fmt.Println("nama terlalu panjang")
		default :
		fmt.Println("nama sudah ok")
	}
}
