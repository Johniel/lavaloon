package main

import (
	"fmt"

	"github.com/johniel/lavaloon/lavaloon"
)

func show(program string) {
	ts, err := lavaloon.Lex(program)
	if err != nil {
		panic(err)
	}

	root, err := lavaloon.Parse(ts)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", root)
}

func main() {
	show("(defun add (a b) (+ a b))")
	show("(defun fact (n) (if (= n 0) 1 (* n (fact (- n 1)))))")
	return
}
