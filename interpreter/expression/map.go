package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Map defines a map expression.
type Map struct {
	interpreter.PosImpl
	Map map[string]interpreter.Expr
}

// Invoke the expression and return a result.
func (expr *Map) Invoke(env *interpreter.Env) (reflect.Value, error) {
	m := make(map[string]interface{})
	for k, e := range expr.Map {
		v, err := e.Invoke(env)
		if err != nil {
			return v, interpreter.NewError(e, err)
		}
		m[k] = v.Interface()
	}
	return reflect.ValueOf(m), nil
}

// Assign a value to the expression and return it.
func (expr *Map) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
