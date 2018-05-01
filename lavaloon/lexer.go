package lavaloon

import (
	"fmt"
	"go/token"
	"regexp"
)

type pattern struct {
	Type   token.Token
	Regexp *regexp.Regexp
}

func patterns() []pattern {
	return []pattern{
		{WHITESPACE, regexp.MustCompile(`^\s+`)},
		{token.COMMENT, regexp.MustCompile(`^;.*`)},
		{token.STRING, regexp.MustCompile(`^("(\\.|[^"])*")`)},
		{token.INT, regexp.MustCompile(`^((([0-9]+)?\.)?[0-9]+)`)},
		{OPEN, regexp.MustCompile(`^(\()`)},
		{CLOSE, regexp.MustCompile(`^(\))`)},
		{SYMBOL, regexp.MustCompile(`^('|[^\s();]+)`)},
	}
}

//
func Lex(program string) ([]*Token, error) {
	ts := make([]*Token, 0)

	for pos := 0; pos < len(program); {
		tmp := pos
		for _, patt := range patterns() {
			if matches := patt.Regexp.FindStringSubmatch(program[pos:]); matches != nil {
				if len(matches) > 1 {
					t := &Token{patt.Type, matches[1], token.Position{"", pos, 0, 0}}
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
