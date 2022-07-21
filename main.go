package main

import (
	"fmt"
	"os"

	"github.com/tommaso-borgato/jakarta-migration/jakarta"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage %s <INPUT_FOLDER>\n", os.Args[0])
		os.Exit(3)
	}
	inputDirectory := os.Args[1]
	fmt.Printf("%s %s\n", os.Args[0], inputDirectory)

	jakarta.ProcessDirectory(inputDirectory)

}
