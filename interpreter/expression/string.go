package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// String defines a string expression.
type String struct {
	interpreter.PosImpl
	Lit string
}

// Invoke the expression and return a result.
func (expr *String) Invoke(env *interpreter.Env) (reflect.Value, error) {
	return reflect.ValueOf(expr.Lit), nil
}

// Assign a value to the expression and return it.
func (expr *String) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
