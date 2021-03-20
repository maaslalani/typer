package wrap

import "bytes"

// Words takes in a long line of text and width as inputs and returns a string
// that contains the contents of `text` with line breaks inserted between words.
// Extremely naive implementation.
func Words(text string, width int) string {
	if len(text) < width {
		return text
	}

	var buffer bytes.Buffer

	count := 1
	for _, c := range text {
		if count >= width && (c == ' ' || c == '\n') {
			buffer.WriteRune('\n')
			count = 0
		} else {
			buffer.WriteRune(c)
		}
		count += 1
	}

	return buffer.String()
}
