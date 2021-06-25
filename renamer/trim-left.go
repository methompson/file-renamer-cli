package renamer

type TrimLeft struct {
	quant int
}

func (tl TrimLeft) Rename(text string, iter int) string {
	return NewTrimBetween(0, tl.quant, Left).Rename(text, iter)
}

func NewTrimLeft(quant int) TrimLeft {
	return TrimLeft{quant}
}
