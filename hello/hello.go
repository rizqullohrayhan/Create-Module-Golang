package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Loger, including
	// the log entry prefix and a flag to disable printing
	// te time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Rizqulloh", "Rayhan", "Ferdiansyah"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	// message, err := greetings.Hello("Rayhan")
	// If an error was returned, print it to te console and
	// exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print thhe returned message
	// to the console
	fmt.Println(messages)
}
