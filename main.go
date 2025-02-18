package main

import (
	"fmt"
	"log"
	"os"

	"ascii-art/utils"
)

// This program takes a command-line argument to generate ASCII art.
func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run . <Text>") // Validate input
	}

	inputstr := os.Args[1]

	result, err := utils.Process(inputstr)

	if err {
		os.Exit(0) // Exit on error
	}

	fmt.Print(result) // Output the generated ASCII art
}
