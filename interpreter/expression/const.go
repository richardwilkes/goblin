package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Const defines a constant.
type Const struct {
	interpreter.PosImpl
	Value string
}

// Invoke the expression and return a result.
func (expr *Const) Invoke(env *interpreter.Env) (reflect.Value, error) {
	switch expr.Value {
	case "true":
		return reflect.ValueOf(true), nil
	case "false":
		return reflect.ValueOf(false), nil
	}
	return reflect.ValueOf(nil), nil
}

// Assign a value to the expression and return it.
func (expr *Const) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
