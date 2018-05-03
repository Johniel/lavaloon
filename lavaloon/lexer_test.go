package lavaloon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Lexer(t *testing.T) {
	assert := assert.New(t)

	var tokens []*Token
	var err error

	tokens, err = Lex("(defun add (a b) (+ a b))")
	assert.Nil(err)
	assert.Len(tokens, 13)

	tokens, err = Lex("(defun fact (n) (if (= n 0) 1 (* n (fact (- n 1)))))")
	assert.Nil(err)
	assert.Len(tokens, 28)

	tokens, err = Lex("'(1 2 3)")
	assert.Nil(err)
	assert.Len(tokens, 6)

	tokens, err = Lex("\"string\"")
	assert.Nil(err)
	assert.Len(tokens, 1)
}
