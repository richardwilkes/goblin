package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Pair defines a map key/value pair expression.
type Pair struct {
	interpreter.PosImpl
	Key   string
	Value interpreter.Expr
}

// Invoke the expression and return a result.
func (expr *Pair) Invoke(env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewStringError(expr, "Not invokable")
}

// Assign a value to the expression and return it.
func (expr *Pair) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
