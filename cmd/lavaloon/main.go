package main

import (
	"os"

	"github.com/johniel/lavaloon/lavaloon"
)

func main() {
	if len(os.Args) != 2 {
		panic("usage: lavaloon FILEPATH")
	}

	filepath := os.Args[1]

	if _, err := os.Stat(filepath); err != nil {
		panic("no such file: " + filepath)
	}

	err := lavaloon.Transpile(filepath)
	if err != nil {
		panic(err)
	}

	return
}
