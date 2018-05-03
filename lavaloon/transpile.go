package lavaloon

import (
	"fmt"
	"go/token"
	"io/ioutil"
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
	fmt.Printf("%v\n", root)
	return nil
}
