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
	pathFlag := flag.String("p", ".", "path root to begin checking files for string, cli args takes precedence over this flag")
	stringFlag := flag.String("s", `"TODO"`, "target string to check files for")
	exactFlag := flag.Bool("e", false, "should program search for exact match (true) or substring (false)")

	flag.Parse()

	// override flag value if filepath was specified as cli args
	argv := flag.Args()
	path := *pathFlag
	if len(argv) > 0 {
		path = argv[0]
	}

	// get absolute filepath
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in getting absolute filepath %v\n", err)
		fmt.Fprintf(os.Stderr, "Falling back to relative filepath\n")
	} else {
		path = absPath
	}

	// select which String Test to check each File with
	var rt readerTest.ReaderTest
	if *exactFlag {
		rt = readerTest.ExactStringTest(*stringFlag)
	} else {
		rt = readerTest.HasSubstringTest(*stringFlag)
	}

	wf := fileContains.NewFileTestWalkFunction(os.Stdout, *stringFlag, rt)
        err = filepath.Walk(path, wf)
        if err != nil {
		fmt.Fprintf(os.Stderr, "Error in running filepath.Walk %v\n", err)
		os.Exit(1)
        }


}
