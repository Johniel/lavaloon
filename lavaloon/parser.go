package lavaloon

import (
	"go/ast"
)

type Parser interface {
	Parse([]*Token) ast.Node
}
