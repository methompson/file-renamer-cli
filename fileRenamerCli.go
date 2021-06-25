package main

import (
	"fmt"
	"os"
)

func main() {
	files := make([]string, 0)

	files = append(files, "123456789", "test2", "TEST3", "Test 1")

	// ops, err := ParseCommandLineArgs(os.Args)

	ops, err := ParseCommandLineArgs([]string{
		"",
		"-lc",
		"-r",
		"test",
		"fun ",
		"-il",
		"good ",
		"0",
		"-ir",
		".ext",
		"0",
		"-r",
		"  ",
		" ",
		"-tc",
		"-il",
		" ",
		"0",
		"-c",
		"0",
		"5",
		"2",
	})

	if err != nil {
		fmt.Println("Invalid Arguments:", err)
		os.Exit(1)
	}

	for iter, file := range files {
		output := file
		for _, op := range ops {
			output = op.Rename(output, iter)
		}
		fmt.Println("old", file)
		fmt.Println("new", output)
	}
}
