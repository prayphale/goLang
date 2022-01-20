package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	english := englishBot{}
	spanish := spanishBot{}
	printGreeting(english)
	printGreeting(spanish)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello World!"
}

func (spanishBot) getGreeting() string {
	return "Hola Mundo"
}
