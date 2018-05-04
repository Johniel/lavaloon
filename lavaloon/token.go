package lavaloon

import (
	"go/token"
)

const (
	WHITESPACE token.Token = iota
	OPEN
	CLOSE
	SYMBOL

	DEFUN
	DO
)

type Token struct {
	Type token.Token
	Val  string
	Pos  token.Position
}
