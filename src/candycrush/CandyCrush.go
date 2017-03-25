/*
	Main program for Candy Crush game
*/
package main

import (
	"fmt"
	"os"
)

// Main function to run Candy Crush game
func main() {
	// Defensive programming
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./candycrush.exe -V | filename")
		os.Exit(1)
	}

	// Get version of game
	if os.Args[1] == "-V" {
		fmt.Println("1.0.0")
		os.Exit(0)
	}

	// Setup model
	SetupModel(os.Args[1])

	// Create the controller/view
	SetupController()

	// Success -> end game
	os.Exit(0)
}
