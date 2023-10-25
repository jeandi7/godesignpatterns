package creational

import (
	"strings"
	"testing"
)

func TestFilesAndFolder(t *testing.T) {

	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{name: "Folder1"}

	folder1.add(file1)

	folder2 := &Folder{name: "Folder2"}

	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	file1Result := file1.search("rose")
	expectedText := "File1(rose)"

	if !strings.Contains(file1Result, expectedText) {
		t.Errorf("must return '%s' not '%s' ", expectedText, file1Result)
	}

	file2Result := folder2.search("rose")
	expectedText2 := "Folder2(rose)File2(rose)File3(rose)Folder1(rose)File1(rose)"

	if !strings.Contains(file2Result, expectedText2) {
		t.Errorf("must return '%s' not '%s' ", expectedText2, file2Result)
	}

}
