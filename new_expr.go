package goblin

import "reflect"

// NewExpr defines a new instance expression.
type NewExpr struct {
	PosImpl
	Type string
}

// Invoke the expression and return a result.
func (expr *NewExpr) Invoke(env *Env) (reflect.Value, error) {
	rt, err := env.Type(expr.Type)
	if err != nil {
		return NilValue, NewError(expr, err)
	}
	return reflect.New(rt), nil
}

// Assign a value to the expression and return it.
func (expr *NewExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	return NilValue, NewInvalidOperationError(expr)
}
