package evaluator

import (
	"github.com/cdarne/monkey/ast"
	"github.com/cdarne/monkey/object"
)

var (
	True  = &object.Boolean{Value: true}
	False = &object.Boolean{Value: false}
	Null  = &object.Null{}
)

func Eval(node ast.Node) object.Object {
	switch n := node.(type) {
	case *ast.Program:
		return evalStatements(n.Statements)
	case *ast.ExpressionStatement:
		return Eval(n.Expression)
	case *ast.PrefixExpression:
		right := Eval(n.Right)
		return evalPrefixExpression(n.Operator, right)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: n.Value}
	case *ast.Boolean:
		return nativeBooleanObject(n.Value)
	}
	return nil
}

func evalStatements(statements []ast.Statement) (result object.Object) {
	for _, s := range statements {
		result = Eval(s)
	}
	return result
}

func nativeBooleanObject(b bool) object.Object {
	if b {
		return True
	}
	return False
}

func evalPrefixExpression(operator string, exp object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(exp)
	case "-":
		return evalMinusOperatorExpression(exp)
	}
	return Null
}

func evalBangOperatorExpression(exp object.Object) object.Object {
	switch exp {
	case True:
		return False
	case False:
		return True
	case Null:
		return True
	default:
		return False
	}
}

func evalMinusOperatorExpression(exp object.Object) object.Object {
	if exp.Type() != object.IntegerObj {
		return Null
	}
	value := exp.(*object.Integer).Value
	return &object.Integer{Value: -value}
}
