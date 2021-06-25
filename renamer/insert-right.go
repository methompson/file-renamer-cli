package renamer

import (
	"strings"
)

type InsertRight struct {
	text string
	pos  int
}

func (i InsertRight) Rename(text string, _ int) string {
	totalLen := len(text)
	if totalLen == 0 {
		return text
	}

	split := strings.Split(text, "")

	left := strings.Join(split[:totalLen-i.pos], "")
	right := strings.Join(split[totalLen-i.pos:], "")

	output := strings.Join([]string{left, i.text, right}, "")
	return output
}

func NewInsertRight(text string, pos int) InsertRight {
	return InsertRight{text, pos}
}
