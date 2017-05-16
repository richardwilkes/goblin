package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// AnonCall defines an anonymous calling expression, e.g. func(){}().
type AnonCall struct {
	goblin.PosImpl
	Expr     goblin.Expr
	SubExprs []goblin.Expr
	VarArg   bool
}

// Invoke the expression and return a result.
func (expr *AnonCall) Invoke(env *goblin.Env) (reflect.Value, error) {
	f, err := expr.Expr.Invoke(env)
	if err != nil {
		return f, goblin.NewError(expr, err)
	}
	if f.Kind() == reflect.Interface {
		f = f.Elem()
	}
	if f.Kind() != reflect.Func {
		return f, goblin.NewStringError(expr, "Unknown function")
	}
	call := &Call{Func: f, SubExprs: expr.SubExprs, VarArg: expr.VarArg}
	return call.Invoke(env)
}

// Assign a value to the expression and return it.
func (expr *AnonCall) Assign(rv reflect.Value, env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.NewInvalidOperationError(expr)
}
