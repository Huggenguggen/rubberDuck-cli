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
	ducks        []Duck  //The ducks
	quitting     bool    // Whether we are quitting
}

// instance of drawing
type Duck struct {
	currX    int // Current X position
	currY    int // Current Y position
	velocity int // Speed of the duck
	ori      int // Orientation (1 for right, -1 for left)
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

	// Prepare the canvas
	canvas := make([][]rune, m.canvasHeight)
	for i := range canvas {
		canvas[i] = make([]rune, m.canvasWidth)
		for j := range canvas[i] {
			canvas[i][j] = ' ' // Initialize with empty spaces
		}
	}

	// Render each duck on the canvas using the shared drawing
	for _, duck := range m.ducks {
		var artToRender string
		if duck.ori == -1 {
			artToRender = reverseASCII(m.drawing.Art, m.drawing.width)
		} else {
			artToRender = m.drawing.Art
		}
		lines := strings.Split(artToRender, "\n")
		for i, line := range lines {
			if duck.currY+i+m.drawing.height >= 0 && duck.currY+i+m.drawing.height < m.canvasHeight {
				for j, char := range line {
					if duck.currX+j+m.drawing.width >= 0 && duck.currX+j+m.drawing.width < m.canvasWidth {
						canvas[duck.currY+i+m.drawing.height][duck.currX+j+m.drawing.width] = char
					}
				}
			}
		}
	}

	// Build the final output from the canvas
	var builder strings.Builder
	for _, row := range canvas {
		builder.WriteString(string(row) + "\n")
	}

	return builder.String()
}

// note that the x and y here are not the middle of the duck,
// it is the top left of the duck
func createDuck(x int, y int, velocity int, ori int) Duck {
	return Duck{
		currX:    x,
		currY:    y,
		velocity: velocity,
		ori:      ori,
	}
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

	//reverseDrawing := reverseASCII(drawing.Art, drawing.width)
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	//f.WriteString(fmt.Sprintf("Drawing width: %d\n", drawing.width))
	//f.WriteString(reverseDrawing)
	defer f.Close()

	// Create a couple of ducks at different positions
	ducks := []Duck{
		createDuck(10, 5, 1, 1),   // Duck 1
		createDuck(30, 11, 1, -1), // Duck 2
	}

	// Create a new Bubbletea program with the model containing ducks and the shared drawing
	p := tea.NewProgram(model{
		drawing: drawing,
		ducks:   ducks,
	})

	// Start the Bubbletea program
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
