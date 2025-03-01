package main

import (
	"fmt"

	"github.com/kkato/go-tutorial/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
