package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// Pair defines a map key/value pair expression.
type Pair struct {
	goblin.PosImpl
	Key   string
	Value goblin.Expr
}

// Invoke the expression and return a result.
func (expr *Pair) Invoke(env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.NewStringError(expr, "Not invokable")
}

// Assign a value to the expression and return it.
func (expr *Pair) Assign(rv reflect.Value, env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.NewInvalidOperationError(expr)
}
