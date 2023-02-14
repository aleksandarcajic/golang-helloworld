package main

import (
	"fmt"
	"log"
	"time"
)

var whatToSay string = "Goodbye, cruel world"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	BirthDate   time.Time
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println(whatToSay)
	fmt.Println(saySomething())

	changeUsingPointer(&whatToSay)
	fmt.Println(whatToSay)

	var user = User{
		FirstName: "Alex",
		LastName:  "Cajic",
	}

	log.Println(user.FirstName)
}

func saySomething() int {
	return 10
}

func changeUsingPointer(s *string) {
	*s = "Hello again!!!"
}
