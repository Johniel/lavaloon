package lavaloon

import (
	"fmt"
	"regexp"
)

const (
	whitespaceToken tokenType = iota
	commentToken
	stringToken
	numberToken
	openToken
	closeToken
	symbolToken
)

// Lexer is the interface than apply lexical analysis
type Lexer interface {
	Run(r string) ([]*Token, error)
}

// Lavaloon struct is lavaloon default lexer
type Lex struct{}

type pattern struct {
	Type   tokenType
	Regexp *regexp.Regexp
}

func patterns() []pattern {
	return []pattern{
		{whitespaceToken, regexp.MustCompile(`^\s+`)},
		{commentToken, regexp.MustCompile(`^;.*`)},
		{stringToken, regexp.MustCompile(`^("(\\.|[^"])*")`)},
		{numberToken, regexp.MustCompile(`^((([0-9]+)?\.)?[0-9]+)`)},
		{openToken, regexp.MustCompile(`^(\()`)},
		{closeToken, regexp.MustCompile(`^(\))`)},
		{symbolToken, regexp.MustCompile(`^('|[^\s();]+)`)},
	}
}

// Run method
func (ll *Lex) Run(program string) ([]*Token, error) {
	ts := make([]*Token, 0)

	for pos := 0; pos < len(program); {
		tmp := pos
		for _, patt := range patterns() {
			if matches := patt.Regexp.FindStringSubmatch(program[pos:]); matches != nil {
				if len(matches) > 1 {
					t := &Token{patt.Type, matches[1], Position{}}
					ts = append(ts, t)
				}
				pos += len(matches[0])
				break
			}
		}
		if tmp == pos {
			return ts, fmt.Errorf("unknown token `%s...`", program[pos:pos+10])
		}
	}

	return ts, nil
}
