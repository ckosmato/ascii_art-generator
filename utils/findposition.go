package utils

// GetPos converts characters in a string to their respective positions in the ASCII table.
func GetPos(str string) []int {
	var letterPositions []int
	for _, pos := range str {
		pos := int(pos) - 32                           // Adjust position based on ASCII offset
		letterPositions = append(letterPositions, pos) // Collect adjusted positions
	}
	return letterPositions // Return the list of letter positions
}
