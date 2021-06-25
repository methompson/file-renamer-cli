package renamer

import "strings"

type Replacer struct {
	oldStr string
	newStr string
}

func (r Replacer) Rename(text string, _ int) string {
	splits := strings.Split(text, r.oldStr)
	return strings.Join(splits, r.newStr)
}

func NewReplacer(oldStr string, newStr string) Replacer {
	return Replacer{oldStr, newStr}
}
