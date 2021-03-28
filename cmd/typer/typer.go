package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
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
	initializeFlags()
	text, err := executeFlags()
	if err != nil {
		log.Fatal(err)
	}

	bar, err := progress.NewModel(progress.WithScaledGradient(blue, purple))
	if err != nil {
		panic(err)
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
		fmt.Fprintf(os.Stdin, "Could not open file %s\n", path)
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

// shuffleWords randomizes word order across the string given
func shuffleWords(s string) (string, error) {
	if len(s) == 0 {
		return "", errors.New("shuffle: empty input")
	}

	rand.Seed(time.Now().UTC().UnixNano())

	shuffled := strings.Fields(s)
	n := len(shuffled)

	if n == 1 {
		return s, errors.New("shuffle: too short to shuffle")
	}

	for i := range shuffled {
		pos := rand.Intn(n - 1)
		shuffled[i], shuffled[pos] = shuffled[pos], shuffled[i]
	}

	return strings.Join(shuffled, " "), nil
}
