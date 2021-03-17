package main

import (
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tjarratt/babble"
)

const (
	blue   = "#4776E6"
	purple = "#8E54E9"
	words  = 5
)

func main() {
	bar, err := progress.NewModel(progress.WithScaledGradient(blue, purple))
	if err != nil {
		panic(err)
	}

	// Randomly generate `words` words for the typing test
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	babbler.Count = words

	text := babbler.Babble()

	program := tea.NewProgram(model{
		progress: bar,
		text:     text,
		start:    time.Now(),
	})

	err = program.Start()
	if err != nil {
		panic(err)
	}
}
