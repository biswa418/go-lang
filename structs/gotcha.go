package main

import "fmt"

func gotcha(){
	mySlice := []string{"Hi", "Hello", "There", "Here"}
	updateSlice(mySlice)

	fmt.Println(mySlice)
}

func updateSlice(s []string){
	s[0] = "Bye"
}