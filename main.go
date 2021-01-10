package main

import (
	"flag"
	"fmt"
	"os"

	"findTodo/fileContains"
)

func main() {
	pathFlag := flag.String("p", ".", "path root to begin checking files for string")
	stringFlag := flag.String("s", `"TODO"`, "target string to check files for")
	exactFlag := flag.Bool("e", true, "should program search for exact match (true) or substring (false)")

	flag.Parse()
	//ok, err := fileContains.Driver("a.in", "todo")
	fmt.Println(*exactFlag)

	ok, err := fileContains.Driver(*pathFlag, *stringFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in Driver %v\n", err)
		os.Exit(1)
	}
	fmt.Println(ok)

}
