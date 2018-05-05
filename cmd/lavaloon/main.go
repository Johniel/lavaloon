package main

import (
	"fmt"
	"os"

	"github.com/johniel/lavaloon/lavaloon"
)

func printUsage() {
	fmt.Println("Usage: lavaloon FILEPATH")
}

func main() {
	if len(os.Args) != 2 {
		printUsage()
		os.Exit(1)
	}

	filepath := os.Args[1]

	if _, err := os.Stat(filepath); err != nil {
		fmt.Println("no such file: " + filepath)
		os.Exit(1)
	}

	err := lavaloon.Transpile(filepath)
	if err != nil {
		panic(err)
	}

	return
}
