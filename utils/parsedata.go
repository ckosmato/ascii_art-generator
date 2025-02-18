package utils

import (
	"bufio"
	"log"
	"os"
)

// Data reads lines from a specified text file and returns them as a slice of strings.
func Data() []string {
	format, err := os.Open("banners/standard.txt")
	if err != nil {
		log.Fatal(err) // Exit on file open error
	}
	defer format.Close() // Ensure the file is closed after reading

	var data []string
	scanner := bufio.NewScanner(format)
	for scanner.Scan() {
		data = append(data, scanner.Text()) // Collect each line of text
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err) // Exit on scanning error
	}
	return data // Return the collected data
}
