package utility

import (
	"os"
	"regexp"
	"strings"

	"github.com/tjarratt/babble"
)

// ReadFile returns the file contents as a string
func ReadFile(path string) (string, error) {
	contents, err := os.ReadFile(path)
	return string(contents), err
}

// RandomWords generates a strings with the specified number of random words using Babble
func RandomWords(n int) string {
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	babbler.Count = n
	return babbler.Babble()
}

// AdjustLength shortens a string if it's word count is greater than n
func AdjustLength(s string, n int) string {
	words := strings.Fields(s)
	if len(words) > n {
		words = words[:n]
	}
	s = strings.Join(words, " ")

	return s
}

// AdjustWhitespace replaces every group of whitespace characters with a single space charracter
func AdjustWhitespace(s string) (string, error) {
	reg, err := regexp.Compile(`\s+`)
	if err != nil {
		return "", err
	}

	s = reg.ReplaceAllString(s, " ")

	if s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}
	return s, nil
}

// RemoveNonAlpha removes all non-alphanumeric characters exept whitespace
func RemoveNonAlpha(s string) (string, error) {
	reg, err := regexp.Compile(`[^a-zA-Z0-9\s]+`)
	if err != nil {
		return "", err
	}

	s = reg.ReplaceAllString(s, "")
	return s, nil
}
