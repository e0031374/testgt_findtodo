package fileContains

import (
	"fmt"
	"os"

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

func Driver(path string, target string) (bool, error) {
	out, err := fileContainsString(path, target)
	if err != nil {
		return false, err
	}

	return out, nil
}
