package main

import "fmt"


// two chatbots - english, spanish
// getGreetings()
// show greetings()

type englishBot struct{}
type spanishBot struct{}



func main(){
	// naive way
	eb := englishBot{}
	// sb := spanishBot{}

	printGreeeting(eb)
	// printGreeeting(sb)

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

func printGreeeting(eb englishBot){
	fmt.Println(eb.getGreeting())
}

// func printGreeeting(sb spanishBot){
// 	fmt.Println(sb.getGreeting())
// }
