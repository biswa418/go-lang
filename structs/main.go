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

	// created the pointer to pass
	jimPointer := &jim
	jimPointer.updateName("Jimothy")
	jim.print()
}


func (p personC) print(){
	fmt.Printf("\n%+v", p)
}


func (pointerToP *personC) updateName(newFirstname string){
	(*pointerToP).firstName = newFirstname
}