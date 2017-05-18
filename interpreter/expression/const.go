package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Const defines a constant.
type Const struct {
	interpreter.PosImpl
	Value reflect.Value
}

func (expr *Const) String() string {
	switch expr.Value {
	case interpreter.TrueValue:
		return "true"
	case interpreter.FalseValue:
		return "false"
	default:
		return "nil"
	}
}

// Invoke the expression and return a result.
func (expr *Const) Invoke(env *interpreter.Env) (reflect.Value, error) {
	return expr.Value, nil
}

// Assign a value to the expression and return it.
func (expr *Const) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
