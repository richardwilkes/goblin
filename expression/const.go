package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin/ast"
)

// Const defines a constant.
type Const struct {
	ast.PosImpl
	Value reflect.Value
}

func (expr *Const) String() string {
	switch expr.Value {
	case ast.TrueValue:
		return "true"
	case ast.FalseValue:
		return "false"
	default:
		return "nil"
	}
}

// Invoke the expression and return a result.
func (expr *Const) Invoke(scope ast.Scope) (reflect.Value, error) {
	return expr.Value, nil
}

// Assign a value to the expression and return it.
func (expr *Const) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
