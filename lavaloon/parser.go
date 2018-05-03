package lavaloon

import (
	"fmt"

	stack "github.com/golang-collections/collections/stack"
)

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
