package main

import (
	"bufio"
	"os"
	"strings"
)

type Drawing struct {
	Art    string // The ASCII art as a string
	width  int    // width should not be negative
	height int    // same for height
	ori    bool   // false is right, true is left
}

func getMiddle(drawing Drawing) (int, int) {
	return (drawing.width / 2), (drawing.height / 2)
}

func fileToString(path string) (Drawing, error) {
	file, err := os.Open(path)
	if err != nil {
		return Drawing{}, err
	}
	defer file.Close()

	var maxWidth int
	var height int
	var builder strings.Builder

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		builder.WriteString(line + "\n")
		height++
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	if err := scanner.Err(); err != nil {
		return Drawing{}, err
	}

	drawing := Drawing{
		Art:    builder.String(),
		width:  maxWidth,
		height: height,
		ori:    true,
	}

	return drawing, nil
}
