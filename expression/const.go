package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// Const defines a constant.
type Const struct {
	goblin.PosImpl
	Value string
}

// Invoke the expression and return a result.
func (expr *Const) Invoke(env *goblin.Env) (reflect.Value, error) {
	switch expr.Value {
	case "true":
		return reflect.ValueOf(true), nil
	case "false":
		return reflect.ValueOf(false), nil
	}
	return reflect.ValueOf(nil), nil
}

// Assign a value to the expression and return it.
func (expr *Const) Assign(rv reflect.Value, env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.NewInvalidOperationError(expr)
}
