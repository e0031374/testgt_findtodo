package fileContains

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"findTodo/readerTest"
)

func fileContainsString(path string, target string, rt readerTest.ReaderTest) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, fmt.Errorf("Error opening file %v ::", err)
	}
	defer f.Close()

	out := rt.Test(f)
	return out, nil
}

func NewFileTestWalkFunction(w io.Writer, target string, rt readerTest.ReaderTest) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("Encountered error while accessing path: %v ::", err)
		}

		if info != nil && info.IsDir() {
			return nil
		}

		toPrint, err := fileContainsString(path, target, rt)
		if err != nil {
			return fmt.Errorf("Encountered error while testing File: %v ::", err)
		}

		if toPrint {
			fmt.Fprintf(w, "%v\n", path)
		}

		return nil
	}
}

func Driver(path string, target string) (bool, error) {
	wf := NewFileTestWalkFunction(os.Stdout, target, readerTest.HasSubstringTest(target))
	err := filepath.Walk(path, wf)
	if err != nil {
		return false, err
	}

	return false, nil
}
