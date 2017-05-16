package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/util"
)

// TernaryOp defines a ternary operator expression.
type TernaryOp struct {
	interpreter.PosImpl
	Expr interpreter.Expr
	LHS  interpreter.Expr
	RHS  interpreter.Expr
}

// Invoke the expression and return a result.
func (expr *TernaryOp) Invoke(env *interpreter.Env) (reflect.Value, error) {
	rv, err := expr.Expr.Invoke(env)
	if err != nil {
		return rv, interpreter.NewError(expr, err)
	}
	var choice interpreter.Expr
	if util.ToBool(rv) {
		choice = expr.LHS
	} else {
		choice = expr.RHS
	}
	rv, err = choice.Invoke(env)
	if err != nil {
		return rv, interpreter.NewError(choice, err)
	}
	return rv, nil
}

// Assign a value to the expression and return it.
func (expr *TernaryOp) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
