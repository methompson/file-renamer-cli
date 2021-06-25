package renamer

import (
	"fmt"
	"strings"
)

type Counter struct {
	pos      int
	start    int
	totalLen int
}

func NewCounter(pos int, start int, totalLen int) Counter {
	return Counter{pos, start, totalLen}
}

func (c Counter) Rename(text string, iter int) string {
	num := iter + c.start
	numTxt := fmt.Sprintf("%d", num)

	txtLen := len(numTxt)
	zeroes := ""

	if txtLen < c.totalLen {
		zeroes = strings.Repeat("0", c.totalLen-txtLen)
	}

	if c.pos == 0 {
		return strings.Join([]string{zeroes, numTxt, text}, "")
	}

	split := strings.Split(text, "")

	left := strings.Join(split[:c.pos], "")
	right := strings.Join(split[c.pos:], "")

	output := strings.Join([]string{left, zeroes, numTxt, right}, "")

	return output
}
