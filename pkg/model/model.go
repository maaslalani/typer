package model

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/guptarohit/asciigraph"
	"github.com/maaslalani/typer/pkg/theme"
)

const (
	width = 60.

	// charsPerWord is the average characters per word used by most typing tests
	// to calculate your WPM score.
	charsPerWord = 5.
)

var (
	wpms []float64
)

type Model struct {
	// Percent is a value from 0 to 1 that represents the current completion of the typing test
	Percent  float64
	Progress *progress.Model
	// Text is the randomly generated text for the user to type
	Text []rune
	// Typed is the text that the user has typed so far
	Typed []rune
	// Start and end are the start and end time of the typing test
	Start time.Time
	// Mistakes is the number of characters that were mistyped by the user
	Mistakes int
	// Score is the user's score calculated by correct characters typed
	Score float64
	// Theme is the current color theme
	Theme *theme.Theme
}

// Init inits the bubbletea model for use
func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) updateProgress() (tea.Model, tea.Cmd) {
	m.Percent = float64(len(m.Typed)) / float64(len(m.Text))
	if m.Percent >= 1.0 {
		return m, tea.Quit
	}
	return m, nil
}

// Update updates the bubbletea model by handling the progress bar update
// and adding typed characters to the state if they are valid typing characters
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Start counting time only after the first keystroke
		if m.Start.IsZero() {
			m.Start = time.Now()
		}

		// User wants to cancel the typing test
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}

		// Deleting characters
		if msg.Type == tea.KeyBackspace && len(m.Typed) > 0 {
			m.Typed = m.Typed[:len(m.Typed)-1]
			return m.updateProgress()
		}

		// Ensure we are adding characters only that we want the user to be able to type
		if msg.Type != tea.KeyRunes {
			return m, nil
		}

		char := msg.Runes[0]
		next := rune(m.Text[len(m.Typed)])

		// To properly account for line wrapping we need to always insert a new line
		// Where the next line starts to not break the user interface, even if the user types a random character
		if next == '\n' {
			m.Typed = append(m.Typed, next)

			// Since we need to perform a line break
			// if the user types a space we should simply ignore it.
			if char == ' ' {
				return m, nil
			}
		}

		m.Typed = append(m.Typed, msg.Runes...)

		if char == next {
			m.Score += 1.
		}

		return m.updateProgress()
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
		if c == rune(m.Text[i]) {
			typed += m.Theme.StringColor(m.Theme.Text.Typed, string(c)).String()
		} else {
			typed += m.Theme.StringColor(m.Theme.Text.Error, string(m.Text[i])).String()
		}
	}

	s := fmt.Sprintf(
		"\n  %s\n\n%s%s",
		m.Progress.View(m.Percent),
		typed,
		m.Theme.StringColor(m.Theme.Text.Untyped, string(remaining)).Faint(),
	)

	var wpm float64
	// Start counting wpm after at least two characters are typed
	if len(m.Typed) > 1 {
		wpm = (m.Score / charsPerWord) / (time.Since(m.Start).Minutes())
	}

	if len(m.Typed) > charsPerWord {
		wpms = append(wpms, wpm)
	}

	wpmsCount := wpms
	if len(wpmsCount) <= 0 {
		wpmsCount = []float64{0}
	}

	graph := asciigraph.Plot(
		wpmsCount,
		asciigraph.Height(m.Theme.Graph.Height),
		asciigraph.Width(m.Progress.Width-5),
		asciigraph.Precision(2),
		asciigraph.SeriesColors(m.Theme.GraphColor()),
	)
	s += fmt.Sprintf("\n\n%s\n\nWPM: %.2f\n", graph, wpm)
	return s
}
