package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin/ast"
)

// Pair defines a map key/value pair expression.
type Pair struct {
	ast.PosImpl
	Key   string
	Value ast.Expr
}

// Invoke the expression and return a result.
func (expr *Pair) Invoke(scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewStringError(expr, "Not invokable")
}

// Assign a value to the expression and return it.
func (expr *Pair) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
