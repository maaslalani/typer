package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tjarratt/babble"
)

const (
	step  = 1.
	width = 60.
)

func main() {
	bar, err := progress.NewModel(progress.WithScaledGradient("#4776E6", "#8E54E9"))
	if err != nil {
		panic(err)
	}

	babbler := babble.NewBabbler()
	babbler.Separator = " "
	babbler.Count = 5

	program := tea.NewProgram(model{
		percent:  0,
		progress: bar,
		text:     babbler.Babble(),
		typed:    "",
		start:    time.Now(),
		end:      time.Time{},
	})

	err = program.Start()
	if err != nil {
		panic(err)
	}
}

type model struct {
	percent  float64
	progress *progress.Model
	text     string
	typed    string
	start    time.Time
	end      time.Time
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
		if msg.Type == tea.KeyBackspace && len(m.typed) > 0 {
			m.typed = m.typed[:len(m.typed)-1]
			m.percent = float64(len(m.typed)) / float64(len(m.text))
		}
		if len(msg.Runes) <= 0 {
			return m, nil
		}
		ascii := int(msg.Runes[0])
		if ascii < 32 || ascii > 126 {
			return m, nil
		}
		m.typed += msg.String()
		m.percent = float64(len(m.typed)) / float64(len(m.text))
		if m.percent >= 1.0 {
			m.percent = 1.0
			m.end = time.Now()
			return m, tea.Quit
		}
		return m, nil

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - 4
		if m.progress.Width > width {
			m.progress.Width = width
		}
		return m, nil

	default:
		return m, nil
	}
}

func (m model) View() string {
	remaining := m.text[len(m.typed):]

	var typed string
	for i, c := range m.typed {
		if byte(c) == m.text[i] {
			typed += string(c)
		} else {
			typed += red(string(c))
		}
	}

	s := fmt.Sprintf(`
  %s

  %s%s

  `,
		m.progress.View(m.percent),
		typed,
		faint(remaining),
	)
	return s
}

func faint(s string) string {
	return "\033[2m" + s + "\033[m"
}

func red(s string) string {
	return "\033[101m" + s + "\033[m"
}
