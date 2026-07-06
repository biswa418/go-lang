package main

import "fmt"


// two chatbots - english, spanish
// getGreetings()
// show greetings()

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}



func main(){
	// interface way
	eb := englishBot{}
	sb := spanishBot{}

	printGreeeting(eb)
	printGreeeting(sb)

}


func printGreeeting(b bot){
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// very specific logic specific to english bot
	return "Hello there!"
}

// can omit the variable name if not used -- see above function - func (englishBot) getGreeeting
func (sb spanishBot) getGreeting() string {
	// very specific to spanish bot
	return "Hola!"
}

// func printGreeeting(eb englishBot){
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeeting(sb spanishBot){
// 	fmt.Println(sb.getGreeting())
// }
