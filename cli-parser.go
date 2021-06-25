package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"methompson.com/fileRenamerCli/renamer"
)

type CLIFlag struct {
	flag      string
	argParams int
	parser    func([]string) (renamer.Renamer, error)
}

// Can Throw errors
func checkAndParseTrimLeftFlag(args []string) (renamer.Renamer, error) {
	argsLength := len(args)

	// current index, plus 1 argument parameter, plus 1 for index length
	if argsLength < 1 {
		return renamer.GetNoOpRenamer(), errors.New("invalid parameters for -tl flag")
	}

	val1 := args[0]

	i, err := strconv.Atoi(val1)

	if err != nil {
		return renamer.GetNoOpRenamer(), err
	}

	return renamer.NewTrimLeft(i), nil
}

// Can Throw errors
func checkAndParseTrimRightFlag(args []string) (renamer.Renamer, error) {
	argsLength := len(args)

	// current index, plus 1 argument parameter, plus 1 for index length
	if argsLength < 1 {
		return renamer.GetNoOpRenamer(), errors.New("invalid parameters for -tr flag")
	}

	val1 := args[0]

	i, err := strconv.Atoi(val1)

	if err != nil {
		return renamer.GetNoOpRenamer(), err
	}

	return renamer.NewTrimRight(i), nil
}

func checkAndParseTrimBetweenFlag(args []string) (renamer.Renamer, error) {
	argsLength := len(args)

	// current index, plus 1 argument parameter, plus 1 for index length
	if argsLength < 3 {
		return renamer.GetNoOpRenamer(), errors.New("invalid parameters for -tb flag")
	}

	val1 := args[0]
	val2 := args[1]
	val3 := strings.ToLower(args[2])

	i1, err1 := strconv.Atoi(val1)
	i2, err2 := strconv.Atoi(val2)

	if err1 != nil {
		return renamer.GetNoOpRenamer(), err1
	} else if err2 != nil {
		return renamer.GetNoOpRenamer(), err2
	} else if val3 != "left" && val3 != "right" {
		return renamer.GetNoOpRenamer(), errors.New("3rd argument must be 'left' or 'right")
	}

	var direction renamer.TrimDirection
	if val3 == "left" {
		direction = renamer.Left
	} else {
		direction = renamer.Right
	}

	return renamer.NewTrimBetween(i1, i2, direction), nil
}

func checkAndParseUpperCaseFlag(_ []string) (renamer.Renamer, error) {
	return renamer.NewUpperCase(), nil
}
func checkAndParseLowerCaseFlag(_ []string) (renamer.Renamer, error) {
	return renamer.NewLowerCase(), nil
}
func checkAndParseTitleCaseFlag(_ []string) (renamer.Renamer, error) {
	return renamer.NewTitleCase(), nil
}

func checkAndParseReplaceFlag(args []string) (renamer.Renamer, error) {
	argsLength := len(args)

	// current index, plus 1 argument parameter, plus 1 for index length
	if argsLength < 2 {
		return renamer.GetNoOpRenamer(), errors.New("invalid parameters for -r flag")
	}

	val1 := args[0]
	val2 := args[1]

	return renamer.NewReplacer(val1, val2), nil
}

func checkAndParseInsertLeftFlag(args []string) (renamer.Renamer, error) {
	argsLength := len(args)

	// current index, plus 1 argument parameter, plus 1 for index length
	if argsLength < 2 {
		return renamer.GetNoOpRenamer(), errors.New("invalid parameters for -il flag")
	}

	val1 := args[0]
	val2 := args[1]

	i2, err2 := strconv.Atoi(val2)

	if err2 != nil {
		return renamer.GetNoOpRenamer(), errors.New("invalid parameters for -il flag")
	}

	return renamer.NewInsertLeft(val1, i2), nil
}

func checkAndParseInsertRightFlag(args []string) (renamer.Renamer, error) {
	argsLength := len(args)

	// current index, plus 1 argument parameter, plus 1 for index length
	if argsLength < 2 {
		return renamer.GetNoOpRenamer(), errors.New("invalid parameters for -il flag")
	}

	val1 := args[0]
	val2 := args[1]

	i2, err2 := strconv.Atoi(val2)

	if err2 != nil {
		return renamer.GetNoOpRenamer(), errors.New("invalid parameters for -il flag")
	}

	return renamer.NewInsertRight(val1, i2), nil
}

func checkAndParseCounterFlag(args []string) (renamer.Renamer, error) {
	argsLength := len(args)

	// current index, plus 1 argument parameter, plus 1 for index length
	if argsLength < 3 {
		return renamer.GetNoOpRenamer(), errors.New("invalid parameters for -c flag")
	}

	val1 := args[0]
	val2 := args[1]
	val3 := args[2]

	i1, err1 := strconv.Atoi(val1)
	i2, err2 := strconv.Atoi(val2)
	i3, err3 := strconv.Atoi(val3)

	if err1 != nil || err2 != nil || err3 != nil {
		return renamer.GetNoOpRenamer(), errors.New("invalid parameters")
	}

	return renamer.NewCounter(i1, i2, i3), nil
}

func ParseCommandLineArgs(args []string) ([]renamer.Renamer, error) {
	flagMap := map[string]CLIFlag{
		"-tl": {"-tl", 1, checkAndParseTrimLeftFlag},    // TrimLeft
		"-tr": {"-tr", 1, checkAndParseTrimRightFlag},   // TrimRight
		"-tb": {"-tb", 3, checkAndParseTrimBetweenFlag}, // TrimBetween
		"-uc": {"-uc", 0, checkAndParseUpperCaseFlag},   // UpperCase
		"-lc": {"-lc", 0, checkAndParseLowerCaseFlag},   // LowerCase
		"-tc": {"-tc", 0, checkAndParseTitleCaseFlag},   // TitleCase
		"-r":  {"-r", 2, checkAndParseReplaceFlag},      // Replacer
		"-il": {"-il", 2, checkAndParseInsertLeftFlag},  // InsertLeft
		"-ir": {"-ir", 2, checkAndParseInsertRightFlag}, // InsertRight
		"-c":  {"-ir", 3, checkAndParseCounterFlag},     // Counter
	}

	ops := make([]renamer.Renamer, 0)
	argsLength := len(args)

	index := 1

	for index < argsLength {
		flagVal, prs := flagMap[args[index]]

		if args[index] == "-i" {
			// 	argQuant, prs := flagMap["-i"]

			// 	if !prs {
			// 		return ops, errors.New("invalid flag: -i")
			// 	}

			// 	slice := args[index+1 : index+1+argQuant]
			// 	fmt.Println(slice)

			index += 2
		} else if prs {
			if argsLength < index+flagVal.argParams+1 {
				err := fmt.Sprintf("invalid parameters for %v", flagVal.flag)
				return ops, errors.New(err)
			}

			slicedArgs := args[index+1 : index+1+flagVal.argParams]
			op, er := flagVal.parser(slicedArgs)

			if er != nil {
				return ops, er
			}

			ops = append(ops, op)

			index += flagVal.argParams + 1
		} else {
			err := fmt.Sprintf("Uknonwn Parameter %v", args[index])
			return ops, errors.New(err)
		}
	}

	return ops, nil
}
