package renamer

import (
	"strings"
)

type TrimDirection int

const (
	Left  TrimDirection = iota
	Right TrimDirection = iota
)

type TrimBetween struct {
	pos       int
	quant     int
	direction TrimDirection
}

func (tb TrimBetween) Rename(text string, _ int) string {
	// If there's no trim, return the text
	if tb.quant <= 0 {
		return text
	}

	split := strings.Split(text, "")
	textLen := len(split)

	// If the trim length is gte the text size, return an empty string
	if tb.quant >= textLen {
		return ""
	}

	// If the position is outside the boundaries of the string, return the text
	if tb.pos >= textLen {
		return text
	}

	var left, right int

	if tb.direction == Left {
		left = tb.pos
		right = tb.pos + tb.quant
	} else {
		left = textLen - (tb.quant + tb.pos)
		right = textLen - tb.pos
	}

	leftTxt := strings.Join(split[:left], "")
	rightTxt := strings.Join(split[right:], "")

	return strings.Join([]string{leftTxt, rightTxt}, "")
}

func NewTrimBetween(pos int, quant int, direction TrimDirection) TrimBetween {
	return TrimBetween{pos, quant, direction}
}
