package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Lexy", lastName: "Coats"}
	fmt.Println(alex)
}
