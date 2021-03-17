package main

// bold returns a string wrapped in the ansi sequence to make text appear bold
func bold(s string) string {
	return "\033[1m" + s + "\033[m"
}

// faint returns a string wrapped in the ansi sequence to make text appear faint
func faint(s string) string {
	return "\033[2m" + s + "\033[m"
}

// red returns a string wrapped in the ansi sequence to make text appear with a bright red background
func red(s string) string {
	return "\033[101m" + s + "\033[m"
}
