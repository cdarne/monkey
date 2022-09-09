package evaluator

import (
	"github.com/cdarne/monkey/ast"
	"github.com/cdarne/monkey/object"
)

func Eval(node ast.Node) object.Object {
	switch n := node.(type) {
	case *ast.Program:
		return evalStatements(n.Statements)
	case *ast.ExpressionStatement:
		return Eval(n.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: n.Value}
	}
	return nil
}

func evalStatements(statements []ast.Statement) (result object.Object) {
	for _, s := range statements {
		result = Eval(s)
	}
	return result
}
