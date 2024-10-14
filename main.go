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
	if x == 0 {
		GoodMorningWorld()
	} else if x == 1 {
		GoodMorningWorld()
	}
	return "good evening world"
}