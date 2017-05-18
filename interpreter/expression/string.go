package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/util"
)

// String defines a string expression.
type String struct {
	interpreter.PosImpl
	Lit      string
	value    reflect.Value
	resolved bool
}

func (expr *String) String() string {
	return util.QuotedString(expr.Lit)
}

// Invoke the expression and return a result.
func (expr *String) Invoke(env *interpreter.Env) (reflect.Value, error) {
	if !expr.resolved {
		expr.resolved = true
		expr.value = reflect.ValueOf(expr.Lit)
	}
	return expr.value, nil
}

// Assign a value to the expression and return it.
func (expr *String) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
