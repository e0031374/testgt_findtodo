package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"findTodo/fileContains"
	"findTodo/readerTest"
)

func main() {
	stringFlag := flag.String("s", `"TODO"`, "target string to check files for")
	exactFlag := flag.Bool("e", false, "should program search for exact match (true) or substring (false)")
	absFlag := flag.Bool("a", false, "program to display absolute path to flagged file (true) or relative path (false)")

	flag.Parse()

	path := processPathInput(flag.Args(), *absFlag)
	// select which String Test to check each File with
	rt := getReaderTest(*exactFlag, *stringFlag)

	wf := fileContains.NewFileTestWalkFunction(os.Stdout, *stringFlag, rt)
	err := filepath.Walk(path, wf)
        if err != nil {
		fmt.Fprintf(os.Stderr, "Error in running filepath.Walk %v\n", err)
		os.Exit(1)
        }


}

func getReaderTest(eFlag bool, searchVal string) readerTest.ReaderTest {
	if eFlag {
		return readerTest.ExactStringTest(searchVal)
	} else {
		return readerTest.HasSubstringTest(searchVal)
	}
}

func processPathInput(argv []string, isAbs bool) string {
	path := "."
	if len(argv) > 0 {
		path = argv[0]
	} else {
		fmt.Println("WARNING: Insufficient argument indicating file path, defaulting to using `.` as root directory\n")
	}

	// get absolute filepath
	if isAbs {
		absPath, err := filepath.Abs(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "WARNING: Error in getting absolute filepath %v\n", err)
			fmt.Fprintf(os.Stderr, "WARNING: Falling back to relative filepath\n\n")
		} else {
			path = absPath
		}
	}
	return path
}
