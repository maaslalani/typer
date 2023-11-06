package flags

import (
	"fmt"
	"math/rand"
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
	Quote         bool
}

// It is not very correct to have Quote in this file, but if it is in typer, it causes an import cycle
// TODO make a new file for Quote or something
type Quote struct {
	Text   string `json:"text"`
	Source string `json:"source"`
	Length int    `json:"length"`
	Id     int    `json:"id"`
}

// formatWords applies formatting based on flags
func (f *Flags) FormatWords(s string) (string, error) {
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

// formatQuote applies formatting based on flags for quotes.
// this method ignores the p and c flags because they are quotes
// returns random quote from struct
func (f *Flags) FormatQuote(q struct {
	Quotes []Quote `json:"quotes"`
}) (string, error) {
	var quotes []string

	if f.Length <= 0 {
		f.Length = DefaultLength
	} else if f.Length > MaxLength {
		f.Length = MaxLength
	}

	// for quotes, f.Length is the max length of the quote because there is no
	// information about words count in json file, so we will filter by letters count
	// and use f.Length as max length
	for _, q := range q.Quotes {
		if q.Length >= f.Length {
			s, err := util.AdjustWhitespace(q.Text)
			if err != nil {
				return "", err
			}
			quotes = append(quotes, s)
		}
	}

	return quotes[rand.Intn(len(quotes))], nil
}
