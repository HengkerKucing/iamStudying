package main

import "fmt"

func sayHelloTo (firstName string, lastName string){
	fmt.Println("halo wak,", firstName, lastName)
}

func main () {
	sayHelloTo("aldi", "prawiratama")
}