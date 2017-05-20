package expression

import (
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
)

// Number defines a number expression.
type Number struct {
	ast.PosImpl
	Value reflect.Value
	Err   error
}

func (expr *Number) String() string {
	switch expr.Value.Kind() {
	case reflect.Float64:
		return fmt.Sprint(expr.Value.Float())
	case reflect.Int64:
		return fmt.Sprint(expr.Value.Int())
	default:
		return "<nil>"
	}
}

// Invoke the expression and return a result.
func (expr *Number) Invoke(scope ast.Scope) (reflect.Value, error) {
	return expr.Value, expr.Err
}

// Assign a value to the expression and return it.
func (expr *Number) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
