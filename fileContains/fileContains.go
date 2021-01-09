package fileContains

import (
	"fmt"
	"os"
	"io"
	"path/filepath"

	"findTodo/readerTest"
)

func fileContainsString(path string, target string) (bool, error) {
	f, err := os.Open(path)
	if err != nil { return false, fmt.Errorf("Error opening file %v ::", err) }
	defer f.Close()


	stringTest := readerTest.NewStringTest(target)
	out := stringTest.Test(f)
	return out, nil
}

//func Driver(path string, target string) (bool, error) {
//	out, err := fileContainsString(path, target)
//	if err != nil {
//		return false, err
//	}
//
//	return out, nil
//}

func NewFileContainsWalkFunction(w io.Writer, target string) func(path string) error {
	return func(path string) error {
		toPrint, err := fileContainsString(path, target)
		if err != nil {
			return fmt.Errorf("Encountered error while testing File: %v ::", err)
		} else {
			if toPrint {
				fmt.Fprintf(w, "%v\n", path)
				//fmt.Fprintf(w, "%v", path)
			}
		}
		return nil
	}
}

func Driver(path string, target string) (bool, error) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error ) error {
		if err != nil {
			return fmt.Errorf("Encountered error while accessing path: %v ::", err)
		}

		toPrint, err := fileContainsString(path, target)
		if err != nil {
			return fmt.Errorf("Encountered error while testing File: %v ::", err)
		} else {
			if toPrint {
				fmt.Println(path)
			}
		}
		return nil
	})
	if err != nil {
		return false, err
	}

	return false, nil
}
