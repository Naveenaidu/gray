package main

import (
	"errors"
	"fmt"
	"log"

	"rsc.io/quote"
)

func Hello_draft(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome, name", name)
	return message, nil
}

func main_draft() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("ray_tracer: ")
	log.SetFlags(0)

	message, err := Hello("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
	fmt.Println(quote.Glass())
}
