package expression

import (
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/util"
)

// TernaryOp defines a ternary operator expression.
type TernaryOp struct {
	interpreter.PosImpl
	Expr  interpreter.Expr
	Left  interpreter.Expr
	Right interpreter.Expr
}

func (expr *TernaryOp) String() string {
	return fmt.Sprintf("%v ? %v : %v", expr.Expr, expr.Left, expr.Right)
}

// Invoke the expression and return a result.
func (expr *TernaryOp) Invoke(env *interpreter.Env) (reflect.Value, error) {
	rv, err := expr.Expr.Invoke(env)
	if err != nil {
		return rv, interpreter.NewError(expr, err)
	}
	var choice interpreter.Expr
	if util.ToBool(rv) {
		choice = expr.Left
	} else {
		choice = expr.Right
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
