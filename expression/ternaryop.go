package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin"
	"github.com/richardwilkes/goblin/util"
)

// TernaryOp defines a ternary operator expression.
type TernaryOp struct {
	goblin.PosImpl
	Expr goblin.Expr
	LHS  goblin.Expr
	RHS  goblin.Expr
}

// Invoke the expression and return a result.
func (expr *TernaryOp) Invoke(env *goblin.Env) (reflect.Value, error) {
	rv, err := expr.Expr.Invoke(env)
	if err != nil {
		return rv, goblin.NewError(expr, err)
	}
	var choice goblin.Expr
	if util.ToBool(rv) {
		choice = expr.LHS
	} else {
		choice = expr.RHS
	}
	rv, err = choice.Invoke(env)
	if err != nil {
		return rv, goblin.NewError(choice, err)
	}
	return rv, nil
}

// Assign a value to the expression and return it.
func (expr *TernaryOp) Assign(rv reflect.Value, env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.NewInvalidOperationError(expr)
}
