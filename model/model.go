package model

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
)

const (
	width = 60.

	// charsPerWord is the average characters per word used by most typing tests
	// to calculate your WPM score.
	charsPerWord = 5.
)

type Model struct {
	// Percent is a value from 0 to 1 that represents the current completion of the typing test
	Percent  float64
	Progress *progress.Model
	// Text is the randomly generated text for the user to type
	Text string
	// Typed is the text that the user has typed so far
	Typed string
	// Start and end are the start and end time of the typing test
	Start time.Time
	End   time.Time
	// Mistakes is the number of characters that were mistyped by the user
	Mistakes int
	// Score is the user's score calculated by correct characters typed
	Score float64
}

// Init inits the bubbletea model for use
func (m Model) Init() tea.Cmd {
	return nil
}

// Update updates the bubbletea model by handling the progress bar update
// and adding typed characters to the state if they are valid typing characters
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// User wants to cancel the typing test
		if msg.Type == tea.KeyCtrlC {
			m.End = time.Now()
			return m, tea.Quit
		}

		// Deleting characters
		if msg.Type == tea.KeyBackspace && len(m.Typed) > 0 {
			m.Typed = m.Typed[:len(m.Typed)-1]
		}

		// Ensure we are adding characters only that we want the user to be able to type
		// and that all the typed characters have a rune width of one.
		if len(msg.Runes) > 0 && int(msg.Runes[0]) > 31 && int(msg.Runes[0]) < 127 {
			m.Typed += msg.String()
		}

		// Update progress bar state
		m.Percent = float64(len(m.Typed)) / float64(len(m.Text))
		if m.Percent >= 1.0 {
			m.End = time.Now()
			return m, tea.Quit
		}

		return m, nil

	case tea.WindowSizeMsg:
		m.Progress.Width = msg.Width - 4
		if m.Progress.Width > width {
			m.Progress.Width = width
		}
		return m, nil

	default:
		return m, nil
	}
}

// View shows the current state of the typing test.
// It displays a progress bar for the progression of the typing test,
// the typed characters (with errors displayed in red) and remaining
// characters to be typed in a faint display
func (m Model) View() string {
	remaining := m.Text[len(m.Typed):]

	var typed string
	for i, c := range m.Typed {
		s := string(c)
		if byte(c) == m.Text[i] {
			typed += s
			m.Score += 1.0
			continue
		}
		typed += termenv.String(s).Background(termenv.ANSIBrightRed).String()
	}

	s := fmt.Sprintf("\n  %s\n\n%s%s", m.Progress.View(m.Percent), typed, termenv.String(remaining).Faint())

	// Display words per minute when finished
	if len(m.Typed) >= len(m.Text) {
		s += termenv.String(fmt.Sprintf("\n\nWPM: %.2f\n", (m.Score/charsPerWord)/(m.End.Sub(m.Start).Minutes()))).Bold().String()
	}
	return s
}
