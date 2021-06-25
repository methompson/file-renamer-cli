package renamer

import (
	"testing"
)

func TestUpperCase(t *testing.T) {
	cap := NewUpperCase()

	input1 := "input"
	output1 := cap.Rename(input1, 1)

	if output1 != "INPUT" {
		t.Fatalf(`output1 = %v. Should be 'INPUT'`, output1)
	}

	input2 := "inpUT"
	output2 := cap.Rename(input2, 1)

	if output2 != "INPUT" {
		t.Fatalf(`output2 = %v. Should be 'INPUT'`, output1)
	}

	input3 := "INPUT"
	output3 := cap.Rename(input3, 1)

	if output3 != "INPUT" {
		t.Fatalf(`output3 = %v. Should be 'INPUT'`, output1)
	}
}

func TestLowerCase(t *testing.T) {
	lc := NewLowerCase()

	input1 := "INPUT"
	output1 := lc.Rename(input1, 0)

	if output1 != "input" {
		t.Fatalf(`output1 = %v. Should be 'input'`, output1)
	}

	input2 := "InpuT"
	output2 := lc.Rename(input2, 0)

	if output2 != "input" {
		t.Fatalf(`output2 = %v. Should be 'input'`, output2)
	}

	input3 := "input"
	output3 := lc.Rename(input3, 0)

	if output3 != "input" {
		t.Fatalf(`output3 = %v. Should be 'input'`, output3)
	}
}

func TestTitleCase(t *testing.T) {
	flc := NewTitleCase()

	input1 := "INPUT"
	output1 := flc.Rename(input1, 0)

	if output1 != "Input" {
		t.Fatalf(`output1 = %v. Should be 'Input'`, output1)
	}

	input2 := "InpuT"
	output2 := flc.Rename(input2, 0)

	if output2 != "Input" {
		t.Fatalf(`output2 = %v. Should be 'Input'`, output2)
	}

	input3 := "input"
	output3 := flc.Rename(input3, 0)

	if output3 != "Input" {
		t.Fatalf(`output3 = %v. Should be 'Input'`, output3)
	}

	input4 := "Input"
	output4 := flc.Rename(input4, 0)

	if output4 != "Input" {
		t.Fatalf(`output4 = %v. Should be 'Input'`, output4)
	}
}
