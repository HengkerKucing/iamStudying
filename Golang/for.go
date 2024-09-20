package main

import "fmt"

func main () {

	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke-", counter)
		counter++
	}

	fmt.Println("selesai")

	for counter1 := 1; counter1 <=10; counter1++ {
		fmt.Println("perulanagan ke-", counter1)
	}

	nama := []string{
		"luqman",
		"aldi",
		"prawiratama",
	}
	for i := 0; i < len(nama); i++ {
		fmt.Println(nama[i])
	}

	for index, name := range nama{
		fmt.Println("index", index, "=", name)
	}

	for _, name := range nama{
		fmt.Println(name)
	}


}