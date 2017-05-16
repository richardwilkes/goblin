package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// New defines a new instance expression.
type New struct {
	interpreter.PosImpl
	Type string
}

// Invoke the expression and return a result.
func (expr *New) Invoke(env *interpreter.Env) (reflect.Value, error) {
	rt, err := env.Type(expr.Type)
	if err != nil {
		return interpreter.NilValue, interpreter.NewError(expr, err)
	}
	return reflect.New(rt), nil
}

// Assign a value to the expression and return it.
func (expr *New) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
