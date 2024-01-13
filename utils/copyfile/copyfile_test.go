package copyfile

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestCopyFileSuccess(t *testing.T) {
	content := []byte("This is a temporary Test File..")

	// Create a temporary source file
	srcFile, err := os.CreateTemp("", "src")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer srcFile.Close()
	defer os.Remove(srcFile.Name())

	// Write content to source file
	if _, err := srcFile.Write(content); err != nil {
		t.Fatalf("Failed to write to source file: %v", err)
	}

	// Test the copyFile function
	dstFileName := "testfile.test"
	if err := copyFile(srcFile.Name(), dstFileName); err != nil {
		t.Errorf("Failed to copy file: %v", err)
	}
	defer os.Remove(dstFileName)

	// Read back the content
	result, err := ioutil.ReadFile(dstFileName)
	if err != nil {
		t.Fatalf("Failed to read destination file: %v", err)
	}

	if !reflect.DeepEqual(content, result) {
		t.Errorf("File content mismatch. Got %s, want %s", result, content)
	}
}

func TestCopyFileFail(t *testing.T) {
	// Test the copyFile function
	dstFileName := "testfile.test"
	srcFileName := "srctestfile.test"
	if err := copyFile(srcFileName, dstFileName); err == nil {
		t.Errorf("No Error Found: %v", err)
	}
	defer os.Remove(dstFileName)
}
