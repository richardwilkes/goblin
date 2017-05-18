package expression

import (
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Make defines a make expression.
type Make struct {
	interpreter.PosImpl
	Type string
}

func (expr *Make) String() string {
	return fmt.Sprintf("make(%s)", expr.Type)
}

// Invoke the expression and return a result.
func (expr *Make) Invoke(env *interpreter.Env) (reflect.Value, error) {
	rt, err := env.Type(expr.Type)
	if err != nil {
		return interpreter.NilValue, interpreter.NewError(expr, err)
	}
	if rt.Kind() == reflect.Map {
		return reflect.MakeMap(reflect.MapOf(rt.Key(), rt.Elem())).Convert(rt), nil
	}
	return reflect.Zero(rt), nil
}

// Assign a value to the expression and return it.
func (expr *Make) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
