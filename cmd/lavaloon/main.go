package main

import (
	"fmt"

	"github.com/johniel/lavaloon/lavaloon"
)

func main() {
	ll := lavaloon.Lex{}
	ts, err := ll.Run("(defun add (a b) (+ a b))")
	if err != nil {
		panic(err)
	}

	for _, t := range ts {
		fmt.Printf("%s\n", t.Val)
	}

	return
}
