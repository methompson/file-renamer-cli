package renamer

import (
	"strings"
)

type InsertLeft struct {
	text string
	pos  int
}

func (i InsertLeft) Rename(text string, _ int) string {
	if len(i.text) == 0 {
		return text
	}

	split := strings.Split(text, "")

	left := strings.Join(split[:i.pos], "")
	right := strings.Join(split[i.pos:], "")

	output := strings.Join([]string{left, i.text, right}, "")
	return output
}

func NewInsertLeft(text string, pos int) InsertLeft {
	return InsertLeft{text, pos}
}
