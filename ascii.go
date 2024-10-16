package main

import (
	"strings"
)

func reverseASCII(ascii_art string, width int) string {
	// Map of characters to their reversed counterparts
	reverseSymbols := map[rune]rune{
		'>':  '<',
		'<':  '>',
		'/':  '\\',
		'\\': '/',
		'[':  ']',
		']':  '[',
		'{':  '}',
		'}':  '{',
		'(':  ')',
		')':  '(',
		'`':  '´',
		'´':  '`',
	}

	// Split the ASCII art into lines
	lines := strings.Split(ascii_art, "\n")

	// Process each line
	for i, line := range lines {
		// Calculate left padding
		leftPadding := width - len(line)

		// Get the content without left padding
		content := strings.TrimLeft(line, " ")
		runes := []rune(content)
		length := len(runes)
		reversed := make([]rune, length)

		// Reverse the line and swap symbols
		for j, char := range runes {
			if reverseChar, exists := reverseSymbols[char]; exists {
				reversed[length-1-j] = reverseChar
			} else {
				reversed[length-1-j] = char
			}
		}

		// Calculate right padding for the reversed line
		rightPadding := width - leftPadding - length
		if rightPadding < 0 {
			rightPadding = 0
		}

		// Create the padded, reversed line
		paddedLine := strings.Repeat(" ", leftPadding) +
			string(reversed) +
			strings.Repeat(" ", rightPadding)

		lines[i] = paddedLine
	}

	// Join the lines back together
	return strings.Join(lines, "\n")
}
