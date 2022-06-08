package flags

import (
	"fmt"
	"strings"

	util "github.com/maaslalani/typer/pkg/utility"
)

const (
	DefaultLength = 20
	MaxLength     = 500
)

type Flags struct {
	Length        int
	MinWordLength int
	Capital       bool
	Punctuation   bool
}

// formatText applies formatting based on flags
func (f *Flags) FormatText(s string) (string, error) {
	var err error
	s, err = util.AdjustWhitespace(s)
	if err != nil {
		return "", err
	}

	if !f.Punctuation {
		s, err = util.RemoveNonAlpha(s)
		if err != nil {
			return "", err
		}
	}

	if f.MinWordLength > 1 {
		s, err = util.MinWordLength(s, f.MinWordLength)
		if err != nil {
			return "", err
		}
	}

	if f.Length <= 0 {
		f.Length = DefaultLength
	} else if f.Length > MaxLength {
		f.Length = MaxLength
	}
	s = util.AdjustLength(s, f.Length)

	if !f.Capital {
		s = strings.ToLower(s)
	}

	if strings.Trim(s, " \n") == "" {
		return s, fmt.Errorf("word list is empty")
	}

	return s, nil
}
