package renamer

type Renamer interface {
	Rename(string, int) string
}

type NoOpRenamer struct{}

func (nop NoOpRenamer) Rename(text string, _ int) string {
	return text
}

func GetNoOpRenamer() NoOpRenamer {
	return NoOpRenamer{}
}
