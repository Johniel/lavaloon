package lavaloon

import (
	"go/ast"
	"go/token"
)

type LavaloonNode struct {
	Token *Token
	Child []*LavaloonNode

	// memoize
	pos token.Pos
	end token.Pos
}

// Pos implements go/ast.Node interface
func (n *LavaloonNode) Pos() token.Pos {
	if n.Token == nil {
		return token.Pos(n.Token.Pos.Offset)
	}
	if n.pos != -1 {
		return n.pos
	}
	n.pos = n.Child[0].Pos()
	return n.pos
}

// End implements go/ast.Node interface
func (n *LavaloonNode) End() token.Pos {
	if n.Token == nil {
		return token.Pos(n.Token.Pos.Offset + len(n.Token.Val))
	}
	if n.end != -1 {
		return n.end
	}
	n.end = n.Child[len(n.Child)-1].End()
	return n.end
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

// Append a child node
func (n *LavaloonNode) Append(m *LavaloonNode) {
	n.Child = append(n.Child, m)
}

func New() *LavaloonNode {
	n := new(LavaloonNode)
	n.pos = -1
	n.end = -1
	return n
}

// Gen method converts *LavaloonNode to *ast.Node
func (n *LavaloonNode) Gen() (*ast.Node, error) {
	return nil, nil
}
