package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// String defines a string expression.
type String struct {
	goblin.PosImpl
	Lit string
}

// Invoke the expression and return a result.
func (expr *String) Invoke(env *goblin.Env) (reflect.Value, error) {
	return reflect.ValueOf(expr.Lit), nil
}

// Assign a value to the expression and return it.
func (expr *String) Assign(rv reflect.Value, env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.NewInvalidOperationError(expr)
}
