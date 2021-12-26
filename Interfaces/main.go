package main

import "fmt"

type bot interface {
	getGreeting() string
}

type english struct{}
type spanish struct{}

func main() {
	english := englishBot{}
	spanish := spanishBot{}
	printGreeting(english)
	printGreeting(spanish)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (english) getGreeting() string {
	return "Hello World!"
}

func (spanish) getGreeting() string {
	return "Hola Mundo"
}
