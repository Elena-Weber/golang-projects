package main

import "fmt"

// type bot is of data type interface
type bot interface {
	// all functions named getGreeting and returning string are of type bot
	getGreeting() string
}

// create separate aliases for type struct (like object in JS)
type englishBot struct{}
type spanishBot struct{}
type russianBot struct{}

func main() {
	// declaring new vars as structs (like object in JS)
	en := englishBot{}
	sp := spanishBot{}
	ru := russianBot{}

	printGreeting(en)
	printGreeting(sp)
	printGreeting(ru)
}

// takes this argument: its shortened 1-letter name and type
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// argument type (receiver), name of function and what it returns
func (englishBot) getGreeting() string {
	return "Hi!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func (russianBot) getGreeting() string {
	return "Privet!"
}