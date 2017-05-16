package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// Paren defines a parent block expression.
type Paren struct {
	goblin.PosImpl
	SubExpr goblin.Expr
}

// Invoke the expression and return a result.
func (expr *Paren) Invoke(env *goblin.Env) (reflect.Value, error) {
	v, err := expr.SubExpr.Invoke(env)
	if err != nil {
		return v, goblin.NewError(expr, err)
	}
	return v, nil
}

// Assign a value to the expression and return it.
func (expr *Paren) Assign(rv reflect.Value, env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.NewInvalidOperationError(expr)
}
