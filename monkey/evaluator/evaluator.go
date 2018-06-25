package evaluator

import (
	"github.com/noissefnoc/book-writing-an-interpreter-in-go/monkey/ast"
	"github.com/noissefnoc/book-writing-an-interpreter-in-go/monkey/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// statement
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	// literal
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}
