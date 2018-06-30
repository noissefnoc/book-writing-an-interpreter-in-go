package evaluator

import (
	"github.com/noissefnoc/book-writing-an-interpreter-in-go/monkey/object"
	"github.com/noissefnoc/book-writing-an-interpreter-in-go/monkey/ast"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
