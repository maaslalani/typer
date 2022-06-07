package flags

import (
	"log"
	"strings"

	util "github.com/maaslalani/typer/pkg/utility"
)

const (
	DefaultLength = 20
	MaxLength     = 500
)

type Flags struct {
	Length      int
	Capital     bool
	Punctuation bool
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

	if f.Length <= 0 {
		log.Println("Length value is incorrect. Restoring to default value.")
		f.Length = DefaultLength
	} else if f.Length > MaxLength {
		log.Println("Max length value exceeded. Restoring to max length value.")
		f.Length = MaxLength
	}
	s = util.AdjustLength(s, f.Length)

	if !f.Capital {
		s = strings.ToLower(s)
	}
	return s, nil
}
