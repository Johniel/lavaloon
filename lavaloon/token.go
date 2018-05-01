package lavaloon

import (
	"go/token"
)

const (
	WHITESPACE token.Token = iota
	OPEN
	CLOSE
	SYMBOL

	DO
)

type Token struct {
	Type     token.Token
	Val      string
	Position token.Position
}
