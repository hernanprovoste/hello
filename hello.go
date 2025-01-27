package main

import (
	"fmt"
	"log"

	"github.com/hernanprovoste/greetings"
)

func main() {
	// Set properties of the logger, including the log entry prefix and a flag to disable printing time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.

	message, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message to the console.

	fmt.Println(message)
}
