package main

import (
	"fmt"
	"go/token"
	"io/ioutil"
	"os"

	"github.com/johniel/lavaloon/lavaloon"
)

func compile(filepath string) error {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	fileset := token.NewFileSet()
	fileset.AddFile(filepath, -1, len(bytes))

	tokens, err := lavaloon.Lex(string(bytes))
	if err != nil {
		return err
	}

	root, err := lavaloon.Parse(tokens)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", root)
	return nil
}

func main() {
	if len(os.Args) != 2 {
		panic("usage: lavaloon FILEPATH")
	}

	filepath := os.Args[1]

	if _, err := os.Stat(filepath); err != nil {
		panic("no such file: " + filepath)
	}

	err := compile(filepath)
	if err != nil {
		panic(err)
	}

	return
}
