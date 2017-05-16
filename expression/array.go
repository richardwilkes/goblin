package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// Array defines an array expression.
type Array struct {
	goblin.PosImpl
	Exprs []goblin.Expr
}

// Invoke the expression and return a result.
func (expr *Array) Invoke(env *goblin.Env) (reflect.Value, error) {
	a := make([]interface{}, len(expr.Exprs))
	for i, e := range expr.Exprs {
		arg, err := e.Invoke(env)
		if err != nil {
			return arg, goblin.NewError(e, err)
		}
		a[i] = arg.Interface()
	}
	return reflect.ValueOf(a), nil
}

// Assign a value to the expression and return it.
func (expr *Array) Assign(rv reflect.Value, env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.NewInvalidOperationError(expr)
}
