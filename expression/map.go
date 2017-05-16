package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// Map defines a map expression.
type Map struct {
	goblin.PosImpl
	Map map[string]goblin.Expr
}

// Invoke the expression and return a result.
func (expr *Map) Invoke(env *goblin.Env) (reflect.Value, error) {
	m := make(map[string]interface{})
	for k, e := range expr.Map {
		v, err := e.Invoke(env)
		if err != nil {
			return v, goblin.NewError(e, err)
		}
		m[k] = v.Interface()
	}
	return reflect.ValueOf(m), nil
}

// Assign a value to the expression and return it.
func (expr *Map) Assign(rv reflect.Value, env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.NewInvalidOperationError(expr)
}
