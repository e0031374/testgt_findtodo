package main

import (
	"fmt"
	"os"
	//"flag"
	
	"findTodo/fileContains"
)

func main() {
	ok, err := fileContains.Driver("a.in", "todo")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in Driver %v\n", err)
		os.Exit(1)
	}
	fmt.Println(ok)

}
