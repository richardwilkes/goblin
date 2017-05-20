package expression

import (
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
)

// Make defines a make expression.
type Make struct {
	ast.PosImpl
	Type string
}

func (expr *Make) String() string {
	return fmt.Sprintf("make(%s)", expr.Type)
}

// Invoke the expression and return a result.
func (expr *Make) Invoke(scope ast.Scope) (reflect.Value, error) {
	rt, err := scope.Type(expr.Type)
	if err != nil {
		return ast.NilValue, ast.NewError(expr, err)
	}
	if rt.Kind() == reflect.Map {
		return reflect.MakeMap(reflect.MapOf(rt.Key(), rt.Elem())).Convert(rt), nil
	}
	return reflect.Zero(rt), nil
}

// Assign a value to the expression and return it.
func (expr *Make) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
