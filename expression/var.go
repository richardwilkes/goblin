package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// Var defines an expression that defines a variable.
type Var struct {
	goblin.PosImpl
	LHS goblin.Expr
	RHS goblin.Expr
}

// Invoke the expression and return a result.
func (expr *Var) Invoke(env *goblin.Env) (reflect.Value, error) {
	rv, err := expr.RHS.Invoke(env)
	if err != nil {
		return rv, goblin.NewError(expr, err)
	}
	if rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}
	return expr.LHS.Assign(rv, env)
}

// Assign a value to the expression and return it.
func (expr *Var) Assign(rv reflect.Value, env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.NewInvalidOperationError(expr)
}
