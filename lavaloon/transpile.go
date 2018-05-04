package lavaloon

import (
	"go/format"
	"go/token"
	"io/ioutil"
	"os"
)

func Transpile(filepath string) error {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	fileset := token.NewFileSet()
	fileset.AddFile(filepath, -1, len(bytes))

	tokens, err := Lex(string(bytes))
	if err != nil {
		return err
	}

	root, err := Parse(tokens)
	if err != nil {
		return err
	}

	generated, err := root.Gen()
	if err != nil {
		return err
	}

	format.Node(os.Stdout, token.NewFileSet(), generated)
	return nil
}
