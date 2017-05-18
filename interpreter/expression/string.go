package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/util"
)

// String defines a string expression.
type String struct {
	interpreter.PosImpl
	Value reflect.Value
}

func (expr *String) String() string {
	return util.QuotedString(expr.Value.String())
}

// Invoke the expression and return a result.
func (expr *String) Invoke(env *interpreter.Env) (reflect.Value, error) {
	return expr.Value, nil
}

// Assign a value to the expression and return it.
func (expr *String) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
