package main

import "fmt"

// all the keys are have to be a singular type -- the values too
func main(){
	// var colors map[string]string
	// colors := make(map[string]string)

	colors := map[string]string{
		"red": "#ff0000",
		"yellow": "#d0ee0e",
	}

	colors["white"] = "#fff" // cannot write colors.white cause if key is int -- it creates issues

	// fmt.Println(colors)
	// delete(colors, "white")
	// fmt.Println(colors)
	delete(colors, "reddg")
	printMap(colors)
}


func printMap(c map[string] string){
	for col, hex := range c {
		fmt.Println("Hex code for", col, "is", hex)
	}
}