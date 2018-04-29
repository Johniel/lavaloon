package lavaloon

import (
	"io"
)

// Lexer is the interface than apply lexical analysis
type Lexer interface {
	Run(r io.ByteReader) ([]*Token, error)
}

// Lavaloon struct is lavaloon default lexer
type LavaloonLexer struct{}

// Run method
func (ll *LavaloonLexer) Run(r io.ByteReader) ([]*Token, error) {
	t := make([]*Token, 0)
	return t, nil
}
