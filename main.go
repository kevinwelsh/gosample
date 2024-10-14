package main

import "fmt"

func main() {
	fmt.Println(HelloWorld())
}

// HelloWorld is a function that returns a string containing "hello world".
func HelloWorld() string {
	return "hello world"
}

// GoodMorningWorld is a function that returns a string containing "good morning world".
func GoodMorningWorld() string {
	return "good morning world"
}


// GoodEveningWorld is a function that returns a string containing "good evening world".
func GoodEveningWorld() string {
	// Sonar non-compliant code:

	return 1 // Noncompliant because non-reachable code below:

	return "good evening world"
}