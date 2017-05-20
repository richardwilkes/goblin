package expression

import (
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
)

// Paren defines a parent block expression.
type Paren struct {
	ast.PosImpl
	SubExpr ast.Expr
}

func (expr *Paren) String() string {
	return fmt.Sprintf("(%v)", expr.SubExpr)
}

// Invoke the expression and return a result.
func (expr *Paren) Invoke(scope ast.Scope) (reflect.Value, error) {
	v, err := expr.SubExpr.Invoke(scope)
	if err != nil {
		return v, ast.NewError(expr, err)
	}
	return v, nil
}

// Assign a value to the expression and return it.
func (expr *Paren) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
