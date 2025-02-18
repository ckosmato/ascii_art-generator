
# ASCII Art Generator in Go

## Overview
This Go program takes an input string from the command line and generates ASCII art representations of the provided text using predefined ASCII characters. The program reads ASCII art mappings from a text file and converts each character in the input text to its corresponding art.

## Features
- **Command-line Input**: Accepts an input string from the command line.
- **ASCII Art Generation**: Converts each character of the input string into ASCII art based on predefined patterns.
- **Supports Multiple Lines**: Processes multiple lines of input by splitting on the newline escape character `\n`.

## Installation
1. Ensure you have **Go** installed.
2. Clone the repository or copy the source file.
3. Build the program:
   ```sh
   go build -o asciiart main.go
   ```

## Usage
The program requires a single command-line argument: the text you want to convert into ASCII art.

```sh
go run . "<Text>"
```

For example:

```sh
go run . "Hello"
```

This will generate ASCII art for the word **Hello** based on the mappings in the `banners/standard.txt` file.

### Important:
- The ASCII art mappings are read from a file named `banners/standard.txt`. Make sure this file exists and contains the correct ASCII art representations.

## Code Explanation

### `main` Function
- Validates the input, ensuring exactly one argument is passed.
- Processes the input string using the `Process` function.
- Outputs the generated ASCII art.

### `Process` Function
- Splits the input string by newline (`\n`).
- Converts each word in the input string into its ASCII art representation by calling `GetPos` to fetch the corresponding positions of characters and mapping them to the ASCII art lines from `banners/standard.txt`.

### `Data` Function
- Reads the ASCII art data from the `banners/standard.txt` file and returns the lines as a slice of strings.

### `GetPos` Function
- Converts each character in the input string into its ASCII value, adjusts the value by subtracting 32 (to map it to printable characters), and returns the positions for the corresponding ASCII art representation.

## Example

### Input (command line):
```sh
go run . "Hello World"
```

### Output (ASCII Art):
```
(Generated ASCII Art for "Hello World" will appear here based on the `standard.txt` mappings)
```

## Error Handling
- If no command-line argument is provided, the program will output a usage message and exit.
- If the ASCII art file `banners/standard.txt` cannot be opened, the program will exit with an error message.

## License
This project is licensed under the MIT License
