package main

import "fmt"

// all the keys are have to be a singular type -- the values too
func main(){
	// var colors map[string]string

	colors := make(map[string]string)

	// colors := map[string]string{
	// 	"red": "#ff0000",
	// 	"yellow": "#d0ee0e",
	// }

	colors["white"] = "#fff"
	colors["red"] = "#ff0000"

	// cannot write colors.white cause if key is int -- it creates issues

	fmt.Println(colors)
	delete(colors, "white")
	fmt.Println(colors)
}