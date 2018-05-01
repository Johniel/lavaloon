package lavaloon

import (
	"fmt"

	"go/token"

	stack "github.com/golang-collections/collections/stack"
)

type LavaloonNode struct {
	Token *Token
	Child []*LavaloonNode
}

// Pos implements go/ast.Node interface
func (n *LavaloonNode) Pos() token.Pos {
	return 0
}

// End implements go/ast.Node interface
func (n *LavaloonNode) End() token.Pos {
	return 0
}

// String implements Stringer interface
func (n *LavaloonNode) String() string {
	if len(n.Child) == 0 {
		return n.Token.Val
	}

	var s string
	s += "("
	for idx, c := range n.Child {
		if idx != 0 {
			s += " "
		}
		s += c.String()
	}
	s += ")"
	return s
}

func (n *LavaloonNode) Append(m *LavaloonNode) {
	n.Child = append(n.Child, m)
}

func New() *LavaloonNode {
	return new(LavaloonNode)
}

func Parse(ts []*Token) (*LavaloonNode, error) {
	s := stack.New()
	s.Push(New())

	for _, t := range ts {
		switch t.Type {
		case OPEN:
			n := New()
			s.Peek().(*LavaloonNode).Append(n)
			s.Push(n)
		case CLOSE:
			s.Pop()
		default:
			n := New()
			n.Token = t
			s.Peek().(*LavaloonNode).Append(n)
		}
	}

	if s.Len() != 1 {
		return nil, fmt.Errorf("invalid token: %v", s.Peek().(*LavaloonNode).Token.Val)
	}
	return s.Peek().(*LavaloonNode), nil
}
