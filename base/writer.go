package base

import "fmt"

type Writer struct {
}

const (
	TEXT_BLACK = iota + 30
	TEXT_RED
	TEXT_GREEN
	TEXT_YELLOW
	TEXT_BLUE
	TEXT_MAGENTE
	TEXT_CYAN
	TEXT_WHITE
)

// Return a string with shell color.
func (w Writer) ColorText(color int, text string, ob ...interface{}) string {
	if color == 0 {
		return text
	}
	return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, text)
}

// Get the color number from string.
func (w Writer) ColorFromString(str string) int {
	var color int
	switch str {
	case "black":
		color = TEXT_BLACK
	case "red":
		color = TEXT_RED
	case "green":
		color = TEXT_GREEN
	case "yellow":
		color = TEXT_YELLOW
	case "blue":
		color = TEXT_BLUE
	case "magenta":
		color = TEXT_MAGENTE
	case "cyan":
		color = TEXT_CYAN
	case "white":
		color = TEXT_WHITE

	}
	return color
}
