package renamer

import (
	"strings"
)

type UpperCase struct{}

func (c UpperCase) Rename(text string, _ int) string {
	return strings.ToUpper(text)
}

func NewUpperCase() UpperCase {
	return UpperCase{}
}

type LowerCase struct{}

func (lc LowerCase) Rename(text string, _ int) string {
	return strings.ToLower(text)
}

func NewLowerCase() LowerCase {
	return LowerCase{}
}

type TitleCase struct{}

func (flc TitleCase) Rename(text string, _ int) string {
	// First convert the text to lower case, then convert all beginnings of wrods to Title
	return strings.Title(strings.ToLower(text))
}

func NewTitleCase() TitleCase {
	return TitleCase{}
}
