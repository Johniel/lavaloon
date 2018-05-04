package lavaloon

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
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
		if n.Token == nil {
			return "()" // e.g. empty function arguments list
		}
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

func (n *LavaloonNode) IsEmpty() bool {
	return n.Token == nil && len(n.Child) == 0
}

func (n *LavaloonNode) IsInternal() bool {
	return !n.IsSymbol()
}

func (n *LavaloonNode) IsLeaf() bool {
	return n.IsSymbol()
}

func (n *LavaloonNode) IsSymbol() bool {
	return n.Token != nil
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

func (n *LavaloonNode) genImport() (*ast.GenDecl, error) {
	if len(n.Child) != 2 {
		return nil, fmt.Errorf("invalid number of arguments(%d): import", len(n.Child))
	}
	if n.Child[0].Token == nil || n.Child[0].Token.Val != "import" {
		panic("")
	}

	return &ast.GenDecl{
		Tok: token.IMPORT,
		Specs: []ast.Spec{
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: n.Child[1].Token.Val,
				},
			},
		},
	}, nil
}

func (n *LavaloonNode) genBlockStmt() (*ast.BlockStmt, error) {
	return &ast.BlockStmt{
		List: []ast.Stmt{
			&ast.ExprStmt{
				X: &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   ast.NewIdent("fmt"),
						Sel: ast.NewIdent("Println"),
					},
					Args: []ast.Expr{
						&ast.BasicLit{
							Kind:  token.STRING,
							Value: strconv.Quote("hello world"),
						},
					},
				},
			},
		},
	}, nil
}

func (n *LavaloonNode) genDefun() (*ast.FuncDecl, error) {
	if len(n.Child) != 4 {
		return nil, fmt.Errorf("invalid number of arguments(%d): defun", len(n.Child))
	}
	if n.Child[0].Token == nil || n.Child[0].Token.Val != "defun" {
		panic("")
	}

	block, err := n.Child[3].genBlockStmt()
	if err != nil {
		return nil, err
	}

	return &ast.FuncDecl{
			Name: ast.NewIdent(n.Child[1].Token.Val),
			Type: &ast.FuncType{},
			Body: block,
		},
		nil
}

func (n *LavaloonNode) Gen() (*ast.File, error) {
	TOP_LEVEL := map[string]bool{
		"defun":   true,
		"import":  true,
		"package": true,
		"const":   true,
	}

	decls := make([]ast.Decl, 0)

	for _, m := range n.Child {
		if !m.IsInternal() {
			return nil, fmt.Errorf("invalid top level expr")
		}
		if !m.Child[0].IsSymbol() {
			return nil, fmt.Errorf("invalid top level function")
		}
		if _, ok := TOP_LEVEL[m.Child[0].Token.Val]; !ok {
			return nil, fmt.Errorf("%s is not acceptable in top level stmt.", m.Child[0].Token.Val)
		}

		switch m.Child[0].Token.Val {
		case "import":
			x, err := m.genImport()
			if err != nil {
				return nil, err
			}
			decls = append(decls, []ast.Decl{x}...)
		case "defun":
			x, err := m.genDefun()
			if err != nil {
				return nil, err
			}
			decls = append(decls, []ast.Decl{x}...)
		case "package":
		case "const":
		}
	}

	return &ast.File{
		Name:  ast.NewIdent("main"),
		Decls: decls,
	}, nil
}
