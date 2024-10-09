package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	canvasWidth  int     // Width of the canvas (terminal)
	canvasHeight int     // Height of the canvas (terminal)
	drawing      Drawing // The ASCII art drawing
	quitting     bool    // Whether we are quitting
}

// width then height
func getCenterPos(m model) (int, int) {
	return (m.canvasWidth / 2), (m.canvasHeight / 2)
}

// init the model, returning a command for bt
func (m model) Init() tea.Cmd {
	return nil
}

// basic update handler
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q": // Exit on ctrl+c or q
			m.quitting = true
			return m, tea.Quit
		}
	case tea.WindowSizeMsg: // Handle window resize
		m.canvasWidth = msg.Width
		m.canvasHeight = msg.Height
		return m, nil
	}
	return m, nil
}

// View renders the canvas and ASCII art centered on the screen.
func (m model) View() string {
	if m.quitting {
		return "Goodbye!\n"
	}

	mX, mY := getCenterPos(m)

	// Prepare the output by building a canvas with empty lines and filling in the ASCII art
	var builder strings.Builder
	lines := strings.Split(m.drawing.Art, "\n")

	for i := 0; i < mY; i++ { // Add empty lines before the ASCII art
		builder.WriteString("\n")
	}

	for _, line := range lines {
		builder.WriteString(strings.Repeat(" ", mX) + line + "\n") // Add x offset to center
	}

	return builder.String()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <path to ascii art file>")
		return
	}

	// Read the ASCII art from the file
	drawing, err := fileToString(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Create a new Bubbletea program with the model
	p := tea.NewProgram(model{drawing: drawing})

	// Start the Bubbletea program
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
