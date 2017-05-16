package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// Make defines a make expression.
type Make struct {
	goblin.PosImpl
	Type string
}

// Invoke the expression and return a result.
func (expr *Make) Invoke(env *goblin.Env) (reflect.Value, error) {
	rt, err := env.Type(expr.Type)
	if err != nil {
		return goblin.NilValue, goblin.NewError(expr, err)
	}
	if rt.Kind() == reflect.Map {
		return reflect.MakeMap(reflect.MapOf(rt.Key(), rt.Elem())).Convert(rt), nil
	}
	return reflect.Zero(rt), nil
}

// Assign a value to the expression and return it.
func (expr *Make) Assign(rv reflect.Value, env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.NewInvalidOperationError(expr)
}
