package expression

import (
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
)

// New defines a new instance expression.
type New struct {
	ast.PosImpl
	Type string
}

func (expr *New) String() string {
	return fmt.Sprintf("new(%s)", expr.Type)
}

// Invoke the expression and return a result.
func (expr *New) Invoke(scope ast.Scope) (reflect.Value, error) {
	rt, err := scope.Type(expr.Type)
	if err != nil {
		return ast.NilValue, ast.NewError(expr, err)
	}
	return reflect.New(rt), nil
}

// Assign a value to the expression and return it.
func (expr *New) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
