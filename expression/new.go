package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// New defines a new instance expression.
type New struct {
	goblin.PosImpl
	Type string
}

// Invoke the expression and return a result.
func (expr *New) Invoke(env *goblin.Env) (reflect.Value, error) {
	rt, err := env.Type(expr.Type)
	if err != nil {
		return goblin.NilValue, goblin.NewError(expr, err)
	}
	return reflect.New(rt), nil
}

// Assign a value to the expression and return it.
func (expr *New) Assign(rv reflect.Value, env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.NewInvalidOperationError(expr)
}
