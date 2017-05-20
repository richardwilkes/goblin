package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin/ast"
	"github.com/richardwilkes/goblin/util"
)

// String defines a string expression.
type String struct {
	ast.PosImpl
	Value reflect.Value
}

func (expr *String) String() string {
	return util.QuotedString(expr.Value.String())
}

// Invoke the expression and return a result.
func (expr *String) Invoke(scope ast.Scope) (reflect.Value, error) {
	return expr.Value, nil
}

// Assign a value to the expression and return it.
func (expr *String) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
