package lavaloon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Parser(t *testing.T) {
	assert := assert.New(t)

	tokens, err := Lex("(defun add (a b) (+ a b))")
	assert.NoError(err)

	root, err := Parse(tokens)
	assert.NoError(err)
	// do
	assert.Len(root.Child, 1)
	assert.Nil(root.Token)
	// defun
	// add
	// (a b)
	// (+ a b)
	assert.Len(root.Child[0].Child, 4)
	assert.Nil(root.Child[0].Token)
	// defun token
	assert.Len(root.Child[0].Child[0].Child, 0)
	assert.Equal(root.Child[0].Child[0].Token.Val, "defun")
}
