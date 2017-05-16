package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Paren defines a parent block expression.
type Paren struct {
	interpreter.PosImpl
	SubExpr interpreter.Expr
}

// Invoke the expression and return a result.
func (expr *Paren) Invoke(env *interpreter.Env) (reflect.Value, error) {
	v, err := expr.SubExpr.Invoke(env)
	if err != nil {
		return v, interpreter.NewError(expr, err)
	}
	return v, nil
}

// Assign a value to the expression and return it.
func (expr *Paren) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
