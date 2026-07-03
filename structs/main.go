package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{firstName: "Lexy", lastName: "Coats"}
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"


	fmt.Println(alex)
	// fmt.Printf("%+v", alex)
	// fmt.Printf("\n%+v", returnpersonC())
	jim := returnpersonC()
	jim.updateName("Jimothy")
	jim.print()
}


func (p personC) print(){
	fmt.Printf("\n%+v", p)
}


func (p personC) updateName(newFirstname string){
	p.firstName = newFirstname
}