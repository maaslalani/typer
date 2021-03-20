package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maaslalani/typer/model"
	wrap "github.com/mitchellh/go-wordwrap"
	"github.com/tjarratt/babble"
)

const (
	blue         = "#4776E6"
	purple       = "#8E54E9"
	words        = 15
	defaultWidth = 60
)

func main() {
	bar, err := progress.NewModel(progress.WithScaledGradient(blue, purple))
	if err != nil {
		panic(err)
	}

	var text string
	if len(os.Args) > 1 {
		text = readFile(os.Args[1])
	} else {
		text = randomWords(words)
	}

	program := tea.NewProgram(model.Model{
		Progress: bar,
		Text:     wrap.WrapString(text, defaultWidth),
		Start:    time.Now(),
	})

	err = program.Start()
	if err != nil {
		panic(err)
	}
}

// readFile returns the file contents as a string
func readFile(path string) string {
	contents, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stdin, "Could not open file %s\n", os.Args[1])
		os.Exit(1)
	}
	return string(contents)
}

// randomWords generates a strings with the specified number of words
func randomWords(n int) string {
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	babbler.Count = n
	return babbler.Babble()
}
