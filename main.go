package main

import (
	"fmt"
	"go-beginner/greetings"
	"log"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Ryan")

	if err != nil {
		log.Fatal(err)
	}

	var array [2]string

	array[0] = "Hello"
	array[1] = "World"

	fmt.Println(array[0], array[1])

	fmt.Println(message)
}
