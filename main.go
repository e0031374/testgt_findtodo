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
	pathFlag := flag.String("p", ".", "path root to begin checking files for string")
	stringFlag := flag.String("s", `"TODO"`, "target string to check files for")
	exactFlag := flag.Bool("e", false, "should program search for exact match (true) or substring (false)")

	flag.Parse()
	fmt.Println(*exactFlag)

	var rt readerTest.ReaderTest
	if *exactFlag {
		rt = readerTest.ExactStringTest(*stringFlag)
	} else {
		rt = readerTest.HasSubstringTest(*stringFlag)
	}

	wf := fileContains.NewFileTestWalkFunction(os.Stdout, *stringFlag, rt)
        err := filepath.Walk(*pathFlag, wf)
        if err != nil {
		fmt.Fprintf(os.Stderr, "Error in running filepath.Walk %v\n", err)
		os.Exit(1)
        }


}
