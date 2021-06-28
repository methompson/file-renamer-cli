package renamer

import (
	"testing"
)

func TestSplitExtension(t *testing.T) {
	var filename string
	var result Filename

	filename = "1.jpg"
	result = SplitExtension(filename)

	if result.Name != "1" || result.Ext != "jpg" {
		t.Fatalf(`result.name = '%v' and result.ext = '%v'. Should be '1' and 'jpg`, result.Name, result.Ext)
	}

	filename = "1"
	result = SplitExtension(filename)

	if result.Name != "1" || result.Ext != "" {
		t.Fatalf(`result.name = '%v' and result.ext = '%v'. Should be '1' and ''`, result.Name, result.Ext)
	}

	filename = "testreallylongname"
	result = SplitExtension(filename)

	if result.Name != "testreallylongname" || result.Ext != "" {
		t.Fatalf(`result.name = '%v' and result.ext = '%v'. Should be 'testreallylongname' and ''`, result.Name, result.Ext)
	}

	filename = "test.really.long.name"
	result = SplitExtension(filename)

	if result.Name != "test.really.long" || result.Ext != "name" {
		t.Fatalf(`result.name = '%v' and result.ext = '%v'. Should be 'test.really.long' and 'name'`, result.Name, result.Ext)
	}

	filename = ".startsWithDot"
	result = SplitExtension(filename)

	if result.Name != ".startsWithDot" || result.Ext != "" {
		t.Fatalf(`result.name = '%v' and result.ext = '%v'. Should be '.startsWithDot' and ''`, result.Name, result.Ext)
	}
}
