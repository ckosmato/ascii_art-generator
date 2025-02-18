package utils

import (
	"strings"
)

// Process converts input text into ASCII art by mapping characters to their corresponding art representations.
func Process(input string) (string, bool) {
	allLines := strings.Split(input, "\\n") // Split input into lines based on newline escape
	result := ""
	lines := Data() // Retrieve ASCII art lines

	for _, word := range allLines {
		if word == "" {
			result += "\n"
			continue
		}
		moreLines := make([]string, 8)
		positions := GetPos(word) // Get the positions of characters in the input

		for _, eachPosition := range positions {
			start := eachPosition*9 + 1 // Calculate starting index for ASCII art
			for i := 0; i < 8; i++ {
				if start+1 < len(lines) {
					moreLines[i] += lines[start+i]
				}
			}
		}
		for _, line := range moreLines {
			result += line + "\n"
		}
	}

	// Handle special cases for the output
	if result == "\n" {
		return result, true
	}

	if result == "\n\n" {
		result = "\n"
		return result, false
	}

	return result, false // Return the final ASCII art result
}
