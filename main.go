package main

import "fmt"

var whatToSay string = "Goodbye, cruel world"

func main() {
	fmt.Println("Hello World!")
	fmt.Println(whatToSay)
	fmt.Println(saySomething())

	changeUsingPointer(&whatToSay)
	fmt.Println(whatToSay)
}

func saySomething() int {
	return 10
}

func changeUsingPointer(s *string) {
	*s = "Hello again!!!"
}