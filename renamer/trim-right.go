package renamer

type TrimRight struct {
	quant int
}

func (tl TrimRight) Rename(text string, iter int) string {
	return NewTrimBetween(0, tl.quant, Right).Rename(text, iter)
}

func NewTrimRight(quant int) TrimRight {
	return TrimRight{quant}
}
