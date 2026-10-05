package main

import (
	"fmt"
	"interpreter-go/01/repl"
	"os"
	"os/user"
)

func main() {
	// Getting the current user of the operating system.
	user, err := user.Current()

	// If we cannot get the current user, stop the program.
	if err != nil {
		panic(err)
	}

	// Showing a greeting message with the current username.
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)

	fmt.Printf("Feel free to type in commands\n")

	// Starting our REPL.
	// os.Stdin → gives the REPL input from the terminal.
	// os.Stdout → gives the REPL output to the terminal.
	repl.Start(os.Stdin, os.Stdout)
}
