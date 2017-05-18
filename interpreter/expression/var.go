package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Var defines an expression that defines a variable.
type Var struct {
	interpreter.PosImpl
	Left  interpreter.Expr
	Right interpreter.Expr
}

// Invoke the expression and return a result.
func (expr *Var) Invoke(env *interpreter.Env) (reflect.Value, error) {
	rv, err := expr.Right.Invoke(env)
	if err != nil {
		return rv, interpreter.NewError(expr, err)
	}
	if rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}
	return expr.Left.Assign(rv, env)
}

// Assign a value to the expression and return it.
func (expr *Var) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
