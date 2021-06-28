package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"methompson.com/fileRenamerCli/renamer"
)

func main() {
	res, err := ParseCommandLineArgs(os.Args)

	// res, err := ParseCommandLineArgs([]string{
	// 	"",
	// 	"-i",
	// 	"/Users/mathewthompson/Desktop/dev/go/fileRenamerCli/testFiles",
	// 	"-id",
	// 	"-lc",
	// 	"-r",
	// 	"test",
	// 	"fun ",
	// 	"-il",
	// 	"good ",
	// 	"0",
	// 	"-ir",
	// 	".ext",
	// 	"0",
	// 	"-r",
	// 	"  ",
	// 	" ",
	// 	"-tc",
	// 	"-il",
	// 	" ",
	// 	"0",
	// 	"-c",
	// 	"0",
	// 	"5",
	// 	"2",
	// })

	if err != nil {
		log.Fatal("Invalid Arguments:", err)
		os.Exit(1)
	}

	if res.directory == "" {
		log.Fatal("No Input Provided")
		os.Exit(1)
	}

	getNewFileNames(res.directory, res)
}

func getNewFileNames(directory string, res ParseResult) {
	dir, dirErr := os.ReadDir(directory)

	if dirErr != nil {
		log.Fatal("Error Reading Directory", res.directory, dirErr)
		os.Exit(1)
	}

	files := make([]string, 0)
	folders := make([]string, 0)

	// files = append(files, "123456789.ext", "test2", "TEST3", "Test 1")

	for _, dirEntry := range dir {
		// var isDir string
		if !dirEntry.IsDir() {
			files = append(files, dirEntry.Name())
		} else if res.includeDirectories {
			folders = append(folders, path.Join(directory, dirEntry.Name()))
		}
	}

	finishedFiles := make([]renamer.RenameEntry, 0)

	for iter, file := range files {
		var output renamer.Filename

		if res.includeExtension {
			output = renamer.Filename{
				Name: file,
				Ext:  "",
			}
		} else {
			output = renamer.SplitExtension(file)
		}

		for _, op := range *res.ops {
			output.Name = op.Rename(output.Name, iter)
		}

		newName := renamer.JoinExtension(output)

		finishedFiles = append(finishedFiles, renamer.RenameEntry{
			OldName: file,
			NewName: newName,
		})
	}

	printNames(finishedFiles, directory)
	if !res.preview {
		renameFiles(finishedFiles, directory)
	}

	for _, folder := range folders {
		getNewFileNames(folder, res)
	}
}

func printNames(files []renamer.RenameEntry, dir string) {
	fmt.Println("Current Directory", dir)
	fmt.Println()

	for _, file := range files {
		fmt.Println("old name:", file.OldName)
		fmt.Println("new name:", file.NewName)
		fmt.Println()
	}
}

func renameFiles(files []renamer.RenameEntry, dir string) {
	for _, file := range files {
		oldPath := path.Join(dir, file.OldName)
		newPath := path.Join(dir, file.NewName)

		e := os.Rename(oldPath, newPath)

		if e != nil {
			log.Fatal("Error Renaming File", e)
		}
	}
}
