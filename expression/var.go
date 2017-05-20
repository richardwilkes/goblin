package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin/ast"
)

// Var defines an expression that defines a variable.
type Var struct {
	ast.PosImpl
	Left  ast.Expr
	Right ast.Expr
}

// Invoke the expression and return a result.
func (expr *Var) Invoke(scope ast.Scope) (reflect.Value, error) {
	rv, err := expr.Right.Invoke(scope)
	if err != nil {
		return rv, ast.NewError(expr, err)
	}
	if rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}
	return expr.Left.Assign(rv, scope)
}

// Assign a value to the expression and return it.
func (expr *Var) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
