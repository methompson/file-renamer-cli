package renamer

import "strings"

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

type Filename struct {
	Name string
	Ext  string
}

func JoinExtension(fileName Filename) string {
	if len(fileName.Ext) == 0 {
		return fileName.Name
	}
	return strings.Join([]string{fileName.Name, fileName.Ext}, ".")
}

func SplitExtension(filename string) Filename {
	split := strings.Split(filename, ".")

	// The first check is if there's no dot in the filename.
	// The second check is if the filename is an empty string, i.e. only one dot and
	// the file name starts with a dog
	if len(split) == 1 || (len(split) == 2 && split[0] == "") {
		return Filename{
			Name: filename,
			Ext:  "",
		}
	}

	nameSection := split[:len(split)-1]
	ext := split[len(split)-1]

	return Filename{
		Name: strings.Join(nameSection, "."),
		Ext:  ext,
	}
}

type RenameEntry struct {
	OldName string
	NewName string
}
