package lavaloon

import (
	"go/ast"
	"go/token"
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

// Gen method converts *LavaloonNode to *ast.Node
func (n *LavaloonNode) Gen() (*ast.Node, error) {
	return nil, nil
}
